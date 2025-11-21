package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	//"github.com/sivagaminathan/InventoryEvent_Service/internal/config"
)

func main() {
	r := chi.NewRouter()

	// Production-friendly middleware
	r.Use(middleware.Logger)                    // Request logging
	r.Use(middleware.Recoverer)                 // Panic recovery
	r.Use(middleware.Timeout(60 * time.Second)) // timeout

	// ---------- ROUTES ----------
	// Gets 500 and logs it in case of panic
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// ---------- SERVER CONFIG ----------
	//cfg := config.LoadConfig()
	//port := fmt.Sprintf("%d", cfg.Port)
	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Graceful shutdown setup
	idleConnsClosed := make(chan struct{})
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		<-sigChan
		log.Println("Shutting down server...")

		// Allow 10 seconds to gracefully close connections
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Failed to gracefully shutdown: %v", err)
		}

		close(idleConnsClosed)
	}()

	fmt.Println("Server running on port:", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %v", err)
	}

	<-idleConnsClosed
	log.Println("Server stopped.")
}
