package main

import (
	"log"

	"github/j1mb0b/gobank/config"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	store, err := NewPostgresStore(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":8080", store)
	server.Run()
}
