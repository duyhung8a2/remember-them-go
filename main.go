package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Init database
	dbFile := "database.db"
	schemaFile := "database_schema.sql"
	db, err := ConnectDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitDB(db, schemaFile)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	pageHandler := NewPageHandler(db)

	r.Mount("/pages", PageRoutes(pageHandler))

	http.ListenAndServe(":3000", r)
}
