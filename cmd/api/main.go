package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"user-mapping/internal/config"
	"user-mapping/internal/container"
	routes "user-mapping/internal/routes"
)

func main() {
	// configure env
	working_directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// fetch env
	env := os.Getenv("APP_ENV")

	configPath := filepath.Join(working_directory, fmt.Sprintf("appsettings.%s.json", env))
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configPath = "appsettings.Development.json"
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	serviceContainer, jwtHelper, err := container.InitializeContainers(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize containers: %v", err)
	}

	routes := routes.RegisterAppRoutes(serviceContainer, jwtHelper)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
