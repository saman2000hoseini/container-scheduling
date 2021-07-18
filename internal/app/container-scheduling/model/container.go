package model

import (
	"github.com/sirupsen/logrus"
)

type Container struct {
	name      string
	state     string
	jobsCount int
	jobsQueue chan Job
}

const idle = "idle"

func NewContainer(name string) *Container {
	container := &Container{name: name, state: idle, jobsQueue: make(chan Job)}
	go container.Execute()

	return container
}

func (c *Container) AddJob(job Job) {
	c.jobsQueue <- job
	c.jobsCount++
}

func (c *Container) decrement() {
	c.jobsCount--
	if c.jobsCount == 0 {
		c.state = idle
	}
}

func (c *Container) Jobs() int {
	return c.jobsCount
}

func (c *Container) Name() string {
	return c.name
}

func (c *Container) State() string {
	return c.state
}

func (c *Container) Execute() {
	for {
		job := <-c.jobsQueue

		c.state = job.Operation

		if err := job.Handle(c.name); err != nil {
			logrus.Errorf("error while executing %s: %s", job.String(), err.Error())
		}

		c.decrement()
	}
}
