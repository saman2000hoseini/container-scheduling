package server

import (
	"context"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/model"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/scheduler"
	"os"
	"os/signal"
	"syscall"

	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/config"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/handler"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const maxJobs = 10

func main(cfg config.Config) {
	os.Mkdir("./results", 0755)
	e := router.New(cfg)

	jobs := make(chan model.Job, maxJobs)

	requestHandler := handler.NewJobHandler(cfg, jobs)
	jobScheduler := scheduler.NewScheduler(jobs)

	e.POST("/request", requestHandler.UserRequest)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := e.Start(cfg.Server.Address); err != nil {
			logrus.Fatalf("failed to start container scheduling server: %s", err.Error())
		}
	}()

	go jobScheduler.Run()

	logrus.Info("container scheduling server started!")

	s := <-sig

	logrus.Infof("signal %s received", s)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.GracefulTimeout)
	defer cancel()

	e.Server.SetKeepAlivesEnabled(false)

	if err := e.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to shutdown container scheduling server: %s", err.Error())
	}
}

// Register registers server command for container-scheduling binary.
func Register(root *cobra.Command, cfg config.Config) {
	runServer := &cobra.Command{
		Use:   "server",
		Short: "server for container scheduling",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runServer)
}
