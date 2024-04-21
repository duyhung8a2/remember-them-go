package main

import (
	"log"
	"net/http"
	"remember_them/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Init database
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	InitDB(db, DBSchemaFile)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	pageHandler := api.NewPageHandler(db)

	r.Mount("/pages", api.PageRoutes(pageHandler))

	http.ListenAndServe(":3000", r)
}
