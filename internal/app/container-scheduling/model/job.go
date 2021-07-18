package model

import "fmt"

type Job struct {
	Id          uint64
	Operation   string
	Source      string
	Destination string
	IsCode      bool
}

func (j Job) String() string {
	return fmt.Sprintf("Id: %d\nOperation: %s\nSource: %s\nDestination: %s", j.Id, j.Operation, j.Source, j.Destination)
}
