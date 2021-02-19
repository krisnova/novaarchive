package cron

import (
	"sync"
	"time"
)

type Job struct {
	Name     string
	ErrCh    chan error
	Duration time.Duration
	Last     *time.Time
	Work     WorkFunc
	M        *sync.Mutex
}

type JobError struct {
	Job *Job
	E   error
}

func NewJob(name string, d time.Duration, f WorkFunc) *Job {
	return &Job{
		Name:     name,
		Duration: d,
		Work:     f,
		M:        &sync.Mutex{},
	}
}

func (j *Job) StartDuration(ch chan *JobError) {
	now := time.Now()
	// Check
	if j.Last != nil && now.Sub(*j.Last) <= j.Duration {
		return
	}

	j.M.Lock()
	go func() {
		jE := &JobError{
			Job: j,
		}
		err := j.Work(j)
		if err != nil {
			jE.E = err
			ch <- jE
		}
		j.Last = &now
		j.M.Unlock()
		ch <- jE
	}()
}
