package schedular

import (
	"fmt"

	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/handler"
	//"github.com/spf13/cobra"
)

func Run() {
	//delivered_job := request.Job_request{}

	fmt.Print("im in schedular")
	for {
		delivered_job := <-handler.Jobs

		fmt.Print("delivered job is ", delivered_job)

	}
}

// func Register(root *cobra.Command) {
// 	runSchedular := &cobra.Command{
// 		Use:   "schedular",
// 		Short: "schedular for container scheduling",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			main()
// 		},
// 	}

// 	root.AddCommand(runSchedular)
// }
