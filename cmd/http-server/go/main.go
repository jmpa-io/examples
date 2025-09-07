package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Setup server.
	port := readEnv("PORT", "8080")
	fmt.Printf("Starting server on port %s...\n", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	// Start server.
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}

// readEnv reads an environment variable and returns its value. If the variable is not set, it returns the provided default value.
func readEnv(key, defaultValue string ) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

