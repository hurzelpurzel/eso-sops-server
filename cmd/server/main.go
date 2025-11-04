package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/hurzelpurzel/eso-sops-server/internal/handlers"
	"github.com/hurzelpurzel/eso-sops-server/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handlers.HealthHandler)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	srv := server.NewServer(addr, mux)

	log.Printf("starting server on %s", addr)
	if err := srv.ListenAndServe(ctx); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
	log.Println("server stopped")
}
