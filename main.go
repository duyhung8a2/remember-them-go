package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"remember_them/api"
	"remember_them/db"
	"remember_them/models"
	"time"

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
	
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// First create a struct or map to hold JSON data

		// Struct
		type Response struct {
			Message string `json: "message"`
			Status  string `json: "status"`
		}
		data := Response{
			Message: "OK",
			Status:  "success",
		}

		// Map
		// data := map[string]string{
		// 	"message": "OK",
		// 	"status":  "success",
		// }

		// Set content header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Marshal the data into JSON format
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// okByte := []byte("OK")
		w.Write(jsonData)
	})

	pageHandler := api.NewPageHandler(bob)
	r.Mount("/pages", api.PageRoutes(pageHandler))

	port := ":3000"
	log.Printf("Server is running on port %s", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}

}
