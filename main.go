package main

import (
	"log"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Init database
	dbFile := "database.db"

	db, err := InitDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	PrintUserBob(db)
}
