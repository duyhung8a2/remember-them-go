package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"remember_them/api"
	"remember_them/db"
	"remember_them/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stephenafamo/bob"
)

func PrintUserUsingTable(db *bob.DB) {
	ctx := context.Background()
	userTable := models.Users
	users, err := userTable.Query(ctx, db).All()
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Username)
	}
}

func main() {
	// Init database
	bob, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer bob.Close()
	db.InitDB(bob)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	pageHandler := api.NewPageHandler(bob)

	r.Mount("/pages", api.PageRoutes(pageHandler))

	http.ListenAndServe(":3000", r)
}
