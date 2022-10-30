package main

import (
	"github.com/robfig/cron/v3"
	"github.com/ting-app/ting-task-nhk-easy/ting"
	"log"
)

func main() {
	c := cron.New()
	c.AddFunc("30 10 * * *", func() {
		err := ting.RunTask()

		if err != nil {
			log.Printf("Run task error %v\n", err)
		}
	})
	c.Start()
}
