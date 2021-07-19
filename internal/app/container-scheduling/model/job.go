package model

import (
	"fmt"
	"github.com/saman2000hoseini/container-scheduling/internal/pkg/operation"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Job struct {
	Id          uint64
	Operation   string
	Source      string
	Destination string
}

func (j Job) String() string {
	return fmt.Sprintf("Id: %d\nOperation: %s\nSource: %s\nDestination: %s", j.Id, j.Operation, j.Source, j.Destination)
}

const (
	defaultDir    = "./results/"
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

	dir := j.Destination

	if dir == "" {
		dir = defaultDir
	} else {
		dir += "/"
	}

	os.Mkdir(dir, 0755)

	if err := setup(container); err != nil {
		return err
	}

	source := getLastSection(j.Source, "/")

	if !operation.IsDefined(j.Operation) {
		op := getLastSection(j.Operation, "/")
		cmd := exec.Command(dockerCommand, copyFile, j.Operation, container+path+op)
		err := cmd.Run()
		if err != nil {
			cleanup(container)
			return err
		}

		cmd = exec.Command(dockerCommand, copyFile, j.Source, container+path+source)
		err = cmd.Run()
		if err != nil {
			cleanup(container)
			return err
		}

		lng := getLastSection(j.Operation, ".")
		if lng == "py" {
			cmd = exec.Command(dockerCommand, execCommand, container, "/app/syncdependencies.sh")

			err = cmd.Run()
			if err != nil {
				cleanup(container)
				return err
			}

			cmd = exec.Command(dockerCommand, execCommand, container, "pip", "install", "-r", "./requirements.txt")

			err = cmd.Run()
			if err != nil {
				cleanup(container)
				return err
			}

			cmd = exec.Command(dockerCommand, execCommand, container, "python3", "./temp/"+op, "./temp/"+source, "./temp/"+op+".txt")

			out, err = cmd.Output()
			if err != nil {
				cleanup(container)
				return err
			}
		} else if lng == "cpp" {
			cmd = exec.Command(dockerCommand, execCommand, container, "g++", "./temp/"+op)
			err = cmd.Run()
			if err != nil {
				cleanup(container)
				return err
			}

			cmd = exec.Command(dockerCommand, execCommand, container, "./a.out", "./temp/"+source, "/app/temp/"+op+".txt")
			out, err = cmd.Output()
			if err != nil {
				cleanup(container)
				return err
			}
		} else if lng == "c" {
			cmd = exec.Command(dockerCommand, execCommand, container, "gcc", "/app/temp/"+op)
			err = cmd.Run()
			if err != nil {
				cleanup(container)
				return err
			}

			cmd = exec.Command(dockerCommand, execCommand, container, "./a.out", "/app/temp/"+source, "/app/temp/"+op+".txt")
			out, err = cmd.Output()
			if err != nil {
				cleanup(container)
				return err
			}
		}

		if err := cleanup(container); err != nil {
			return err
		}

		return ioutil.WriteFile(dir+op+".out", out, 0644)
	}

	cmd := exec.Command(dockerCommand, copyFile, j.Source, container+path+source)
	logrus.Info(cmd.String())

	err := cmd.Run()
	if err != nil {
		logrus.Errorf("error while transfering input to container: %s", err.Error())

		cleanup(container)
		return err
	}

	cmd = exec.Command(dockerCommand, execCommand, container, j.Operation, "/app/temp/"+source, "/app/temp/"+j.Operation+".txt")
	logrus.Info(cmd.String())

	out, err = cmd.Output()
	if err != nil {
		cleanup(container)
		return err
	}

	if err := cleanup(container); err != nil {
		logrus.Errorf("error while cleaning up container: %s", err.Error())
	}

	return ioutil.WriteFile(dir+j.Operation+".txt", out, 0644)
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
