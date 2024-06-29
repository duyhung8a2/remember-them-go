package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// Load env from .env
	if err := k.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
		l.Fatalf("Error loading .env file: %v", err)
	}
	//Load env from system
	if err := k.Load(env.Provider("", ".", nil), nil); err != nil {
		l.Fatalf("Error loading environment variables: %v", err)
	}

	port := ":" + k.String("SERVER_PORT")
	log.Printf("Server is running on port %s", port)

	r := chi.NewRouter()

	s := &http.Server{
		Addr:         port,
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Receive terminate, graceful shutdown", sig)

	// Wait requests to be finished then shutdown
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
