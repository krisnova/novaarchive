package main

import (
	"time"

	"github.com/kris-nova/logger"
	"github.com/kris-nova/novaarchive/cron"
)

func main() {
	//
	f := func(j *cron.Job) error {
		logger.Always("Running job: %s", j.Name)
		return nil
	}
	job1 := cron.NewJob("Every 3 seconds", time.Second*3, f)
	job2 := cron.NewJob("Every 10 seconds", time.Second*10, f)
	service := cron.NewService("My Service")
	service.Add(job1)
	service.Add(job2)
	jErr := service.Start()
	for {
		j := <-jErr
		logger.Warning(j.E.Error())
	}
}
