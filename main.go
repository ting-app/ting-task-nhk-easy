package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/go-co-op/gocron"
	"github.com/ting-app/ting-task-nhk-easy/ting"
	"log"
	"os"
	"time"
)

func main() {
	enableSentry := os.Getenv("ENABLE_SENTRY") == "true"
	sentryDsn := os.Getenv("SENTRY_DSN")
	runMode := os.Getenv("RUN_MODE")

	if enableSentry {
		if sentryDsn == "" {
			log.Fatal("sentryDsn is required")
		}

		err := sentry.Init(sentry.ClientOptions{
			Dsn:              sentryDsn,
			TracesSampleRate: 1.0,
		})

		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		} else {
			log.Println("sentry enabled")
		}
	}

	if runMode == "immediately" {
		run(enableSentry)()
	} else {
		scheduler := gocron.NewScheduler(time.UTC)
		_, err := scheduler.Every(1).Day().At("10:30").Do(run(enableSentry))

		if err != nil {
			log.Fatalf("Failed to schedule task, %v", err)
		}

		log.Println("Task scheduled")

		scheduler.StartBlocking()
	}
}

func run(enableSentry bool) func() {
	return func() {
		err := ting.RunTask()

		if err != nil {
			log.Printf("Run task error %v\n", err)

			if enableSentry {
				sentry.CaptureException(err)
			}
		}
	}
}
