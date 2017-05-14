package boot

import (
	"os"

	"github.com/branch-app/branch-mono-go/clients/auth"
	"github.com/branch-app/branch-mono-go/libraries/configloader"
	"github.com/branch-app/branch-mono-go/services/halo4/app"
	"github.com/branch-app/branch-mono-go/services/halo4/models"
	"github.com/branch-app/branch-mono-go/services/halo4/server"
	"github.com/branch-app/branch-mono-go/services/halo4/services/waypoint"
)

func RunDebug() {
	// Load config
	var config models.Config
	err := configloader.Unmarshal("debug", &config)
	if err != nil {
		panic(err)
	}

	// Set port from env, otherwise default
	addr := os.Getenv("ADDR")
	if addr == "" && config.Addr != "" {
		addr = config.Addr
	} else if addr == "" {
		addr = "localhost:3000"
	}

	auth := auth.NewClient(config.AuthURL, config.AuthURL)
	waypointService := waypoint.NewClient(auth, config.Mongo, "branch_service-halo4")

	app := app.NewApp(auth, waypointService)
	server := server.NewServer(app)
	server.Run(addr)
}
