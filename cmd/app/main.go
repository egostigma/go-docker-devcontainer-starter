package main

import (
	"log"

	"go-docker-devcontainer-starter/config"
	"go-docker-devcontainer-starter/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
