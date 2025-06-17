package main

import (
	"context"
	"log"

	"github.com/AntiB-Projects/agentic_go/config"
	"github.com/AntiB-Projects/agentic_go/database"
)

func main() {

	// Initialize a context for database operations
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// This single call sets up the connection and all repositories.
	db, err := database.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	log.Println("Database connection successful.")

	usr, err := db.Users.CreateUser(ctx, "test@user.com")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	log.Printf("User created with ID: %d", usr.ID)
}
