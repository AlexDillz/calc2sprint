package main

import (
	"log"

	"github.com/AlexDillz/calc2sprint/internal/config"
	"github.com/AlexDillz/calc2sprint/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	srv := server.NewServer(cfg)

	addr := cfg.ServerAddr + ":" + cfg.Port
	log.Printf("Starting server on %s", addr)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
