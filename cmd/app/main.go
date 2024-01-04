package app

import (
	"github.com/leslieleung/communicator/internal/route"
	"github.com/leslieleung/communicator/pkg/config"
	"github.com/leslieleung/communicator/pkg/database"
	"github.com/leslieleung/communicator/pkg/log"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use: "serve",
	Run: runServe,
}

func runServe(cmd *cobra.Command, args []string) {
	log.Init()
	config.Init(
		config.WithAutomaticEnv(),
		config.WithEnvPrefix("communicator"),
		config.WithDefaultValues(map[string]interface{}{
			"port": 8080,
		}),
	)
	database.Init(config.GetConfig().GetString("dsn"))
	route.StartRouter()
}

func init() {
}
