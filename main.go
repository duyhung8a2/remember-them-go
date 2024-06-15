package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"remember_them/handlers"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

// // Init database
// bob, err := db.ConnectDB()
// if err != nil {
// 	log.Fatal(err)
// }
// defer bob.Close()
// db.InitDB(bob)

// r := chi.NewRouter()

// r.Use(middleware.RequestID)
// r.Use(middleware.RealIP)
// r.Use(middleware.Logger)
// r.Use(middleware.Recoverer)
// r.Use(middleware.Timeout(60 * time.Second))

// pageHandler := api.NewPageHandler(bob)
// r.Mount("/pages", api.PageRoutes(pageHandler))

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(httprate.Limit(10, 1*time.Minute,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Too many request, please try again later",
			})
		})))

	port := ":3000"
	log.Printf("Server is running on port %s", port)

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)
	r.Mount("/", ph.Routes())

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
