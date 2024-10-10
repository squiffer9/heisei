package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"heisei/internal/server/api/handlers"
	"heisei/internal/server/api/middleware"
	"heisei/internal/server/config"
	"heisei/internal/server/repositories"
	"heisei/internal/server/services"
	"heisei/pkg/database"
	"heisei/pkg/utils"

	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	utils.InitLogger(cfg.Log.Level)
	logger := utils.GetLogger()

	// Initialize database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}
	defer db.Close()

	// Run database migrations
	sqlDB, err := db.DB.DB()
	if err != nil {
		logger.Error("Failed to get sql.DB instance", zap.Error(err))
		os.Exit(1)
	}
	if err := database.RunMigrations(sqlDB); err != nil {
		logger.Error("Failed to run database migrations", zap.Error(err))
		os.Exit(1)
	}

	// Initialize repositories and services
	repos := repositories.NewRepositories(db.DB)
	services := services.NewServices(repos, logger)

	// Initialize handlers and middleware
	h := handlers.NewHandlers(services, logger)
	m := middleware.NewMiddleware(logger)

	// Set up HTTP server
	mux := h.SetupRoutes()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: m.LoggingMiddleware(mux),
	}

	// Start server
	go func() {
		logger.Info("Starting server", zap.String("address", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed unexpectedly", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Server is shutting down...")

	// Gracefully shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exiting")
}
