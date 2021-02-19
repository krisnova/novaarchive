package cron

import (
	"sync"
	"time"
)

type Service struct {
	M    map[string]sync.Mutex
	Jobs []*Job
}

func NewService(name string) *Service {
	return &Service{}
}

func (s *Service) Add(j *Job) {
	s.Jobs = append(s.Jobs, j)
}

func (s *Service) Start() chan *JobError {
	ch := make(chan *JobError)
	go func() {
		for {
			// ------ [ Service ] -------
			n := len(s.Jobs)
			for i := 0; i < n; i++ {
				// TODO Nóva to think about how she wants to manage this
				s.Jobs[i].StartDuration(ch)
				time.Sleep(time.Second * 1)
			}
			// ------ [ Service ] -------
		}
	}()
	return ch
}
