package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stephenafamo/bob/dialect/sqlite"
	"github.com/stephenafamo/bob/dialect/sqlite/sm"
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

	queryString, args, err := sqlite.Select(
		sm.From("users"),
	).Build()
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Query(queryString, args...)
	if err != nil {
		log.Fatal(err)
	}

	var users []User

	for result.Next() {
		var user User
		err := result.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	fmt.Println(users)
}
