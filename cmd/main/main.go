package main

import (
	"log"
	"task-microservices/config"
	"task-microservices/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}

	app.Run(cfg)
}
