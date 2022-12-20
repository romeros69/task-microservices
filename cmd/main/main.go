package main

import (
	"log"
	"task-microservices/config"
	"task-microservices/internal/app"
)

// @tittle Task microservice
// @version 1.0
// @description API Server for Task microservice

// @host localhost:9000
// @BasePath /
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}

	app.Run(cfg)
}
