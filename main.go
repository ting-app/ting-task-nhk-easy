package main

import (
	"github.com/go-co-op/gocron"
	"github.com/ting-app/ting-task-nhk-easy/ting"
	"log"
	"time"
)

func main() {
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(1).Day().At("10:30").Do(func() {
		err := ting.RunTask()

		if err != nil {
			log.Printf("Run task error %v\n", err)
		}
	})

	if err != nil {
		log.Fatalf("Failed to schedule task, %v", err)
	}

	log.Println("Task scheduled")

	scheduler.StartBlocking()
}
