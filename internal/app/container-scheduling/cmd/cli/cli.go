package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/model"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/scheduler"
	"github.com/spf13/cobra"

	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/config"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/handler"
	"github.com/sirupsen/logrus"
)

const (
	maxJobs        = 100
	containerName  = "container_scheduler"
	containerCount = 3
)

func main(cfg config.Config) {
	os.Mkdir("./results", 0755)

	jobs := make(chan model.Job, maxJobs)

	containers := make([]*model.Container, containerCount)
	for i := 1; i <= containerCount; i++ {
		containers[i-1] = model.NewContainer(fmt.Sprintf("%s%d", containerName, i))
	}

	requestHandler := handler.CLINewJobHandler(cfg, jobs)
	jobScheduler := scheduler.NewScheduler(jobs, containers)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	logrus.Info("container scheduling started!")

	go jobScheduler.Run()

	reader := bufio.NewReader(os.Stdin)
	for {
		logrus.Info("Enter New Request: ")
		text, _ := reader.ReadString('\n')
		requestHandler.UserRequest(text)
	}
}

func Register(root *cobra.Command, cfg config.Config) {
	runCLI := &cobra.Command{
		Use:   "cli",
		Short: "cli for container scheduling",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runCLI)
}
