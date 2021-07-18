package scheduler

import (
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os/exec"
	"strings"
)

const (
	dir           = "./results/"
	dockerCommand = "docker"
	copyFile      = "cp"
	execCommand   = "exec"
	path          = ":/app/"
)

var defaultContainer = "container_scheduler1"

func handleJob(job model.Job, container string) error {
	var out []byte

	if container == "" {
		container = defaultContainer
	}

	source := getLastSection(job.Source, "/")

	if job.IsCode {
		operation := getLastSection(job.Operation, "/")
		cmd := exec.Command(dockerCommand, copyFile, job.Operation, container+path+operation)
		_, err := cmd.Output()
		if err != nil {
			return err
		}

		cmd = exec.Command(dockerCommand, copyFile, job.Source, container+path+source)
		_, err = cmd.Output()
		if err != nil {
			return err
		}

		op := getLastSection(job.Operation, ".")
		if op == "py" {
			cmd = exec.Command(dockerCommand, execCommand, container, "python3", operation, source)
			out, err = cmd.Output()
			if err != nil {
				return err
			}
		} else if op == "cpp" {
			cmd = exec.Command(dockerCommand, execCommand, container, "g++", operation)
			_, err = cmd.Output()
			if err != nil {
				return err
			}

			cmd = exec.Command(dockerCommand, execCommand, container, "./a", source)
			out, err = cmd.Output()
			if err != nil {
				return err
			}
		} else if op == "c" {
			cmd = exec.Command(dockerCommand, execCommand, container, "gcc", operation)
			_, err = cmd.Output()
			if err != nil {
				return err
			}

			cmd = exec.Command(dockerCommand, execCommand, container, "./a", source)
			out, err = cmd.Output()
			if err != nil {
				return err
			}
		}

		return ioutil.WriteFile(dir+job.Operation+".out", out, 0644)
	} else {
		cmd := exec.Command(dockerCommand, copyFile, job.Source, container+path+source)
		logrus.Info(cmd.String())

		_, err := cmd.Output()
		if err != nil {
			logrus.Errorf("error while transfering input to container: %s", err.Error())
			return err
		}

		cmd = exec.Command(dockerCommand, execCommand, container, job.Operation, source, job.Destination)
		logrus.Info(cmd.String())

		out, err = cmd.Output()
		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(dir+job.Destination, out, 0644)
}

func getLastSection(entry, delimiter string) string {
	separated := strings.Split(entry, delimiter)
	return separated[len(separated)-1]
}
