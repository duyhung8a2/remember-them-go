package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"remember_them/models"
	"strings"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/sqlite"
	"github.com/stephenafamo/bob/dialect/sqlite/sm"
	_ "modernc.org/sqlite"
)

func ConnectDB(dbFile string) (*bob.DB, error) {
	db, err := bob.Open(DBDriver, dbFile)
	if err != nil {
		return nil, err
	}
	return &db, nil
}

func InitDB(db *bob.DB, sqlFileString string) {
	content, err := os.ReadFile(sqlFileString)
	if err != nil {
		log.Fatal(err)
	}

	statements := strings.Split(string(content), ";\n")

	ctx := context.Background()
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := db.ExecContext(ctx, stmt)
		if err != nil {
			log.Fatal(err)
		}
	}
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

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	fmt.Println(users)
}

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
