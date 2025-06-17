package main

import (
	"log"

	"github.com/AntiB-Projects/agentic_go/config"
	"github.com/AntiB-Projects/agentic_go/database"
)

func main() {

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Printf("Loaded configuration: %+v", cfg)

	if err := database.Init(cfg); err != nil {
		log.Fatalf("db init failed: %v", err)
	}
	defer database.DB.Close()
	log.Println("Database connection established successfully")
}
