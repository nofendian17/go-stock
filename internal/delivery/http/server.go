package http

import (
	"context"
	"errors"
	"fmt"
	"go-stock/internal/app"
	"log"
	"net/http"
	"time"
)

func Start(ctx context.Context, bootstrap app.Bootstrap) {
	router := http.NewServeMux()
	RegisterRoutes(router, bootstrap.GetHandler())

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", bootstrap.GetConfig().GetApplication().Host, bootstrap.GetConfig().GetApplication().Port),
		Handler: router,
	}

	// Start the server in a separate goroutine
	go func() {
		log.Println("Starting server on", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for context cancellation
	<-ctx.Done()
	log.Println("Shutting down HTTP server...")

	// Create a timeout context for shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	} else {
		log.Println("HTTP server shut down cleanly")
	}
}
