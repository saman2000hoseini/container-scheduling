package cmd

import (
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/cmd/server"
	"github.com/saman2000hoseini/container-scheduling/internal/app/container-scheduling/config"
	"github.com/spf13/cobra"
)

// NewRootCommand creates a new container-scheduling root command.
func NewRootCommand() *cobra.Command {
	var root = &cobra.Command{
		Use: "container-scheduling",
	}

	cfg := config.Init()

	server.Register(root, cfg)

	//schedular.Register(root)

	return root
}
