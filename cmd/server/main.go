package main

import (
	"log"

	"github.com/donidarmawan/test-go/internal/app"
	"github.com/donidarmawan/test-go/internal/config"
	"github.com/donidarmawan/test-go/internal/database"
)

func main() {
	// 1. Load configuration
	cfg := config.Load()

	// 2. Connect to database & run migrations
	db := database.Connect(cfg)

	// 3. Wire dependencies & register routes
	application := app.New(db, cfg)

	// 4. Start server
	addr := ":" + cfg.AppPort
	log.Printf("Server starting on %s", addr)
	if err := application.Router.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
