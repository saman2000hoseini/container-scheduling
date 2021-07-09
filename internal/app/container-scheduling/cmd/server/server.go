package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/cmd/schedular"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/config"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/handler"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	e := router.New(cfg)

	request_handler := &handler.Job_handler{Cfg: cfg}

	go schedular.Run()

	e.POST("/request", request_handler.User_request)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := e.Start(cfg.Server.Address); err != nil {
			logrus.Fatalf("failed to start virtual-box management server: %s", err.Error())
		}
	}()

	logrus.Info("virtual-box management server started!")

	s := <-sig

	logrus.Infof("signal %s received", s)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.GracefulTimeout)
	defer cancel()

	e.Server.SetKeepAlivesEnabled(false)

	if err := e.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to shutdown virtual-box management server: %s", err.Error())
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
