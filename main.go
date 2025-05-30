package main

import (
	"context"
	"go-stock/internal/app"
	"go-stock/internal/config"
	"go-stock/internal/delivery/cron"
	"go-stock/internal/delivery/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load the configuration
	cfg, err := config.NewConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the application
	bootstrap, err := app.NewBootstrap(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Set up context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the cron scheduler
	go cron.Start(ctx, bootstrap)

	// Start the HTTP server
	go http.Start(ctx, bootstrap)

	// Handle OS shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	log.Printf("Received signal: %s. Initiating shutdown...", sig)

	// Trigger graceful shutdown
	cancel()

	log.Println("Shutdown complete.")
}
