package main

import (
	"log"
	"time"

	"github.com/mileusna/crontab"
)

func main() {
	cron := crontab.New()

	err := cron.AddJob("* * * * *", func() {
		log.Println("Este mensaje automatico se muestra cada un minuto")
	})

	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Minute)
}
