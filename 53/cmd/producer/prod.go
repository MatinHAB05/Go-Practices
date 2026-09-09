package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"test/tasks"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	lport := os.Getenv("REDIS_HOST_PORT")
	pass := os.Getenv("RDB_PASSWORD")
	redisAddr := fmt.Sprintf("localhost:%s", lport)

	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Username: "default",
		Password: pass,
	})
	defer client.Close()

	// ۱. ارسال تسک فوری به صف critical
	emailTask, err := tasks.NewEmailDeliveryTask(42, "welcome_email")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(emailTask, asynq.Queue("critical"))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// ۲. زمان‌بندی تسک برای ۲۴ ساعت بعد در صف default
	info, err = client.Enqueue(emailTask, asynq.ProcessIn(24*time.Hour), asynq.Queue("default"))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("scheduled task: id=%s queue=%s", info.ID, info.Queue)

	// ۳. ارسال تسک سنگین به صف low با آپشن‌های اختصاصی
	imageTask, err := tasks.NewImageResizeTask("https://example.com/image.jpg")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err = client.Enqueue(
		imageTask,
		asynq.Queue("low"),
		asynq.MaxRetry(10),
		asynq.Timeout(3*time.Minute),
	)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task with options: id=%s queue=%s", info.ID, info.Queue)
}
