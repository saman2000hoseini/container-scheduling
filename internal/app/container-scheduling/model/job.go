package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os/exec"
	"strings"
)

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

const (
	dir           = "./results/"
	dockerCommand = "docker"
	copyFile      = "cp"
	execCommand   = "exec"
	path          = ":/app/temp/"
)

var defaultContainer = "container_scheduler1"

func (j Job) Handle(container string) error {
	var out []byte

	if container == "" {
		container = defaultContainer
	}

	if err := setup(container); err != nil {
		return err
	}

	source := getLastSection(j.Source, "/")

	if j.IsCode {
		operation := getLastSection(j.Operation, "/")
		cmd := exec.Command(dockerCommand, copyFile, j.Operation, container+path+operation)
		_, err := cmd.Output()
		if err != nil {
			cleanup(container)
			return err
		}

		cmd = exec.Command(dockerCommand, copyFile, j.Source, container+path+source)
		_, err = cmd.Output()
		if err != nil {
			cleanup(container)
			return err
		}

		op := getLastSection(j.Operation, ".")
		if op == "py" {
			cmd = exec.Command(dockerCommand, execCommand, container, "python3", "./temp/"+operation, "./temp/"+source)
			out, err = cmd.Output()
			if err != nil {
				cleanup(container)
				return err
			}
		} else if op == "cpp" {
			cmd = exec.Command(dockerCommand, execCommand, container, "g++", "./temp/"+operation)
			_, err = cmd.Output()
			if err != nil {
				cleanup(container)
				return err
			}

			cmd = exec.Command(dockerCommand, execCommand, container, "./a", "./temp/"+source)
			out, err = cmd.Output()
			if err != nil {
				cleanup(container)
				return err
			}
		} else if op == "c" {
			cmd = exec.Command(dockerCommand, execCommand, container, "gcc", "./temp/"+operation)
			_, err = cmd.Output()
			if err != nil {
				cleanup(container)
				return err
			}

			cmd = exec.Command(dockerCommand, execCommand, container, "./a", "./temp/"+source)
			out, err = cmd.Output()
			if err != nil {
				cleanup(container)
				return err
			}
		}

		if err := cleanup(container); err != nil {
			return err
		}

		return ioutil.WriteFile(dir+j.Operation+".out", out, 0644)
	}

	cmd := exec.Command(dockerCommand, copyFile, j.Source, container+path+source)
	logrus.Info(cmd.String())

	_, err := cmd.Output()
	if err != nil {
		logrus.Errorf("error while transfering input to container: %s", err.Error())

		cleanup(container)
		return err
	}

	cmd = exec.Command(dockerCommand, execCommand, container, j.Operation, "/app/temp/"+source, "/app/temp/"+j.Destination)
	logrus.Info(cmd.String())

	out, err = cmd.Output()
	if err != nil {
		cleanup(container)
		return err
	}

	if err := cleanup(container); err != nil {
		return err
	}

	return ioutil.WriteFile(dir+j.Destination, out, 0644)
}

func getLastSection(entry, delimiter string) string {
	separated := strings.Split(entry, delimiter)
	return separated[len(separated)-1]
}

func setup(container string) error {
	return exec.Command(dockerCommand, execCommand, container, "mkdir", "temp").Run()
}

func cleanup(container string) error {
	return exec.Command(dockerCommand, execCommand, container, "rm", "-R", "./temp").Run()
}
