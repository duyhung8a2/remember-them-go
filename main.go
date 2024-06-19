package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"remember_them/handlers"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	redocMiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// // Init database
// bob, err := db.ConnectDB()
// if err != nil {
// 	log.Fatal(err)
// }
// defer bob.Close()
// db.InitDB(bob)

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

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Compress(5))
	r.Use(chiMiddleware.Timeout(60 * time.Second))
	r.Use(httprate.Limit(10, 1*time.Minute,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Too many request, please try again later",
			})
		})))

	// Index handler
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	ph := handlers.NewProducts(l)
	r.Mount("/products", ph.Routes())

	options := redocMiddleware.RedocOpts{SpecURL: "/swagger.yaml"}
	r.Mount("/docs", redocMiddleware.Redoc(options, nil))

	workDir, _ := os.Getwd()
	r.Mount("/swagger.yaml", http.FileServer(http.Dir(filepath.Join(workDir, "./"))))

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
