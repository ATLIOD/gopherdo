package main

import (
	"gopherdo/internal/config"
	"gopherdo/internal/db"
	"gopherdo/internal/handler"
	"gopherdo/internal/update"
	"log"
)

var (
	Version   = "dev"
	BuildTime = "unknown"
)

func main() {

	log.Printf("GopherDo Version: %s (Built: %s)", Version, BuildTime)
	cfg := config.NewConfig()

	// Initialize database
	db, err := db.NewDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize handler
	h := handler.NewHandler(db)

	// Check for updates
    release, err := update.CheckForUpdates(Version)
    if err != nil {
        log.Printf("Failed to check for updates: %v", err)
    } else if release != nil {
        log.Printf("New version available: %s\nDownload at: %s",
            release.Version, release.URL)
    }

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := h.Start(cfg.Port); err != nil {
		log.Fatal(err)
	}
}
