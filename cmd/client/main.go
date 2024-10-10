package main

import (
	"log"

	"heisei/internal/client/api"
	"heisei/internal/client/config"
	"heisei/internal/client/tui"
	"heisei/pkg/utils"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger := utils.GetLogger()

	// Initialize API client
	apiClient := api.NewClient(cfg.Server.URL)

	// Initialize and run TUI application
	app, err := tui.NewApp(cfg, apiClient, logger)
	if err != nil {
		logger.Fatal("Failed to initialize TUI", err)
	}

	if err := app.Run(); err != nil {
		logger.Fatal("Application error", err)
	}
}
