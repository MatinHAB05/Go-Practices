package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalln(err)
	}

	// تنظیم زمان برای ساعت ۱۹:۱۰ امروز
	now := time.Now()
	runTime := time.Date(now.Year(), now.Month(), now.Day(), 19, 8, 0, 0, time.Local)

	if runTime.Before(now) {
		runTime = runTime.Add(24 * time.Hour)
	}

	j, err := s.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(runTime),
		),
		gocron.NewTask(
			func(a string, b int) {
				fmt.Println(b, a)
			},
			"hello",
			1,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Job ID:", j.ID())
	fmt.Println("Scheduled for:", runTime)

	s.Start()

	select {
	case <-time.After(time.Minute * 5):
	}

	_ = s.Shutdown()
}
