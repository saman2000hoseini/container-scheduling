package server

import (
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/config"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	//e := router.New(cfg)
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
