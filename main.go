package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"remember_them/handlers"
	"time"
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
	port := ":3000"
	log.Printf("Server is running on port %s", port)

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         port,
		Handler:      sm,
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

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Receive terminate, graceful shutdown", sig)

	s.ListenAndServe()

	// Wait requests to be finished then shutdown
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
