package schedular

import (
	"fmt"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	//"github.com/spf13/cobra"
)

func main() {
	//delivered_job := request.Job_request{}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			deliveredJob := <-handler.Jobs

			fmt.Print("delivered job is ", deliveredJob)
		}
	}()

	logrus.Info("container scheduling scheduler started!")

	s := <-sig

	logrus.Infof("signal %s received", s)
}

func Register(root *cobra.Command) {
	runScheduler := &cobra.Command{
		Use:   "scheduler",
		Short: "scheduler for container scheduling",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}

	root.AddCommand(runScheduler)
}
