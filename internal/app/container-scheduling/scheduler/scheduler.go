package scheduler

import (
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/model"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	jobs chan model.Job
}

func NewScheduler(jobs chan model.Job) *Scheduler {
	return &Scheduler{jobs: jobs}
}

func (s *Scheduler) Run() {
	for {
		deliveredJob := <-s.jobs

		logrus.Infof("New job recieved; %v", deliveredJob)
	}
}
