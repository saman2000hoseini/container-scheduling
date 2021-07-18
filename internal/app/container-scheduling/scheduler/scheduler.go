package scheduler

import (
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/model"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	jobs       chan model.Job
	containers []*model.Container
}

func NewScheduler(jobs chan model.Job, containers []*model.Container) *Scheduler {
	return &Scheduler{jobs: jobs, containers: containers}
}

func (s *Scheduler) Run() {
	for {
		deliveredJob := <-s.jobs

		logrus.Infof("New job recieved:\n%v", deliveredJob)

		minIndex := 0
		for i := range s.containers {
			if s.containers[i].Jobs() == 0 {
				minIndex = i
				break
			}

			if s.containers[i].Jobs() < s.containers[minIndex].Jobs() {
				minIndex = i
			}
		}

		s.containers[minIndex].AddJob(deliveredJob)
	}
}
