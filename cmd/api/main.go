package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"user-mapping/internal/config"
	"user-mapping/internal/container"
	routes "user-mapping/internal/routes"
)

func main() {
	// Get working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Base config path
	baseConfigPath := filepath.Join(workingDirectory, "appsettings.json")

	// Load configuration (LoadConfig will handle APP_ENV internally)
	cfg, err := config.LoadConfig(baseConfigPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize containers
	serviceContainer, jwtHelper, err := container.InitializeContainers(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize containers: %v", err)
	}

	// Register routes
	router := routes.RegisterAppRoutes(serviceContainer, jwtHelper)

	// Use PORT from env first, fallback to config, then default
	port := os.Getenv("PORT")
	if port == "" {
		if cfg.Port != "" {
			port = cfg.Port
		} else {
			port = "5001"
		}
	}

	// Create and start HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server started on port %s in %s environment", port, cfg.AppEnv)
	log.Fatal(srv.ListenAndServe())
}
