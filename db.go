package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/sqlite/sm"
	"github.com/stephenafamo/bob/dialect/sqlite"
	_ "modernc.org/sqlite"
)

func InitDB(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, err
	}

	bobExec := bob.New(db)
	ctx := context.Background()

	// Create a table
	_, err = bobExec.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	)`)
	if err != nil {
		return nil, err
	}

	// _, err = db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
	// if err != nil {
	// 	return nil, err
	// }
	// _, err = db.Exec("INSERT INTO users (name) VALUES (?)", "Bob")
	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}

func PrintUserRaw(db *sql.DB) {
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

func PrintUserBob(db *sql.DB) {
	queryString, args, err := sqlite.Select(
		sm.From("users"),
	).Build()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(queryString, args...)
	if err != nil {
		log.Fatal(err)
	}

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	fmt.Println(users)
}
