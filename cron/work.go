package cron

type WorkFunc func(j *Job) error
