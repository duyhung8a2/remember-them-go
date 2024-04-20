package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	dbFile := "database.db"
	db, err := ConnectDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	InitDB(db)
	PrintUserUsingTable(db)

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

func PageRoutes(pageHandler *PageHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", pageHandler.ListPages)
	r.Post("/", pageHandler.CreatePage)
	r.Get("/{id}", pageHandler.GetPages)
	r.Put("/{id}", pageHandler.UpdatePage)
	r.Delete("/{id}", pageHandler.DeletePage)
	return r
}
