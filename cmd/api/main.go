package main

import (
	"log"
	"net/http"
	"user-mapping/internal/config"
	"user-mapping/internal/container"
	router "user-mapping/internal/register_routes/router"
)

func main() {
	cfg, err := config.LoadConfig("../../appsettings.json")
	if err != nil {
		log.Fatal(err)
	}

	serviceContainer, jwtHelper, err := container.InitializeContainers(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize containers: %v", err)
	}

	routes := router.RegisterAppRoutes(serviceContainer, jwtHelper)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
