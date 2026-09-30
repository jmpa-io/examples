package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Setup server.
	port := readEnv("PORT", "8080")
	slog.Info("Starting server", "port", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	// Listen for shutdown signals.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start server in background.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server error", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for signal then gracefully shut down.
	<-ctx.Done()
	slog.Info("Shutting down server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("Shutdown error", "error", err)
		os.Exit(1)
	}
	slog.Info("Server stopped")
}

// readEnv reads an environment variable and returns its value. If the variable is not set, it returns the provided default value.
func readEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
