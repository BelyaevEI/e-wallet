package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BelyaevEI/e-wallet/internal/app"
)

func main() {

	// Init application
	app, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	// Creating channel for graceful shutdown
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		log.Println("Server is start")
		if err := app.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe error: %v", err)
		}
	}()

	// Given signal for shutdown
	sig := <-sigint
	log.Printf("Received signal: %v", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server
	if err := app.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

}
