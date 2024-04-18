package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Init database
	dbFile := "database.db"

	db, err := InitDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "Bob")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users: ")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %d, Name: %s\n", id, name)
	}

}
