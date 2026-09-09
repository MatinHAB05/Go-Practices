package main

import (
	"fmt"
	"log"
	"os"

	"test/tasks"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	lport := os.Getenv("REDIS_HOST_PORT")
	pass := os.Getenv("RDB_PASSWORD")
	redisAddr := fmt.Sprintf("localhost:%s", lport)

	// ۱. تنظیمات کامل اتصال Redis طبق .env
	redisOpt := asynq.RedisClientOpt{
		Addr:     redisAddr,
		Username: "default",
		Password: pass,
	}

	// ۲. تنظیمات ورکر
	srv := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6, // وزن اولویت بالا
				"default":  3, // وزن اولویت معمولی
				"low":      1, // وزن اولویت پایین
			},
		},
	)

	// ۳. ثبت مسیرهای تسک
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDelivery, tasks.HandleEmailDeliveryTask)
	mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor())

	log.Println("Worker running and listening for tasks...")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
