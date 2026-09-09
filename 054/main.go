package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/riverqueue/river/rivermigrate"
)

// ---------------------------------------------------------
// ۱. تعریف انواع جاب‌ها (Job Args & Workers)
// ---------------------------------------------------------

type EmailArgs struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (EmailArgs) Kind() string { return "send_email" }

type EmailWorker struct {
	river.WorkerDefaults[EmailArgs]
}

func (w *EmailWorker) Work(ctx context.Context, job *river.Job[EmailArgs]) error {
	fmt.Printf("[%s] ⏳ Processing Email to %s...\n", time.Now().Format("15:04:05"), job.Args.To)
	time.Sleep(2 * time.Second)
	fmt.Printf("[%s] ✅ Email Sent to %s!\n", time.Now().Format("15:04:05"), job.Args.To)
	return nil
}

// جاب دوره‌ای ۱: پاکسازی دیتابیس (بر اساس اینتروال)
type IntervalCleanupArgs struct{}

func (IntervalCleanupArgs) Kind() string { return "interval_cleanup" }

type IntervalCleanupWorker struct {
	river.WorkerDefaults[IntervalCleanupArgs]
}

func (w *IntervalCleanupWorker) Work(ctx context.Context, job *river.Job[IntervalCleanupArgs]) error {
	fmt.Printf("[%s] 🔄 [PERIODIC INTERVAL] Running periodic cleanup task...\n", time.Now().Format("15:04:05"))
	time.Sleep(3 * time.Second)
	return nil
}

// جاب دوره‌ای ۲: گزارش‌گیری روزانه (بر اساس کرون جاب)
type DailyReportArgs struct{}

func (DailyReportArgs) Kind() string { return "daily_report" }

type DailyReportWorker struct {
	river.WorkerDefaults[DailyReportArgs]
}

func (w *DailyReportWorker) Work(ctx context.Context, job *river.Job[DailyReportArgs]) error {
	fmt.Printf("[%s] 📊 [CRON JOB] Generating daily analytics report...\n", time.Now().Format("15:04:05"))
	time.Sleep(4 * time.Second)
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, reading environment variables directly")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbPort, dbName)

	ctx := context.Background()

	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbPool.Close()

	migrator, err := rivermigrate.New(riverpgxv5.New(dbPool), nil)
	if err != nil {
		log.Fatalf("Error creating migrator: %v", err)
	}
	_, err = migrator.Migrate(ctx, rivermigrate.DirectionUp, nil)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	workers := river.NewWorkers()
	river.AddWorker(workers, &EmailWorker{})
	river.AddWorker(workers, &IntervalCleanupWorker{})
	river.AddWorker(workers, &DailyReportWorker{})

	// ---------------------------------------------------------
	// ۲. تنظیم جاب‌های دوره‌ای (Periodic / Scheduled Jobs)
	// ---------------------------------------------------------

	// حالت اول: اجرا در بازه‌های زمانی مشخص (مثلاً هر ۱۵ ثانیه)
	intervalJob := river.NewPeriodicJob(
		river.PeriodicInterval(15*time.Second),
		func() (river.JobArgs, *river.InsertOpts) {
			return IntervalCleanupArgs{}, nil
		},
		&river.PeriodicJobOpts{RunOnStart: true},
	)

	// حالت دوم: اجرا با استفاده از فرمت Cron
	// برای فرمت Cron 5-part یا 6-part از river.PeriodicCron استفاده می‌شود
	// cronJob := river.NewPeriodicJob(
	// 	river.PeriodicCron("0 8 * * *", false), // مثلاً هر روز ساعت ۸ صبح
	// 	func() (river.JobArgs, *river.InsertOpts) {
	// 		return DailyReportArgs{}, nil
	// 	},
	// 	&river.PeriodicJobOpts{RunOnStart: false},
	// )

	// ۳. ساخت کلاینت River با جاب‌های دوره‌ای
	riverClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 5},
		},
		Workers:      workers,
		PeriodicJobs: []*river.PeriodicJob{intervalJob /*, cronJob*/},
	})
	if err != nil {
		log.Fatalf("Error creating River client: %v", err)
	}

	if err := riverClient.Start(ctx); err != nil {
		log.Fatalf("Error starting River client: %v", err)
	}

	fmt.Printf("🚀 App Started at %s\n\n", time.Now().Format("15:04:05"))

	// ---------------------------------------------------------
	// ۴. تست سایر گزینه‌های زمان‌بندی (جاب‌های تک‌نوبتی)
	// ---------------------------------------------------------

	// جاب ۱: ارسال فوری
	_, _ = riverClient.Insert(ctx, EmailArgs{
		To:      "instant@example.com",
		Subject: "ایمیل فوری",
	}, nil)

	// جاب ۲: ارسال با ۱۰ ثانیه تاخیر
	_, _ = riverClient.Insert(ctx, EmailArgs{
		To:      "delayed@example.com",
		Subject: "ایمیل ۱۰ ثانیه تاخیر",
	}, &river.InsertOpts{
		ScheduledAt: time.Now().Add(10 * time.Second),
	})

	// جاب ۳: زمان‌بندی روی یک Date و Time دقیق در آینده
	// برای نمونه: تاریخ فردا ساعت ۰۸:۰۰ صبح
	specificDate := time.Date(2026, time.September, 9, 8, 0, 0, 0, time.Local)
	_, _ = riverClient.Insert(ctx, EmailArgs{
		To:      "future@example.com",
		Subject: "ایمیل فردا صبح",
	}, &river.InsertOpts{
		ScheduledAt: specificDate,
	})

	// منتظر ماندن برای Graceful Shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nShutting down gracefully...")
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = riverClient.Stop(ctxTimeout)
}
