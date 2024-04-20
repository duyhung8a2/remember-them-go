package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"remember_them/models"

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

func InitDB(db *bob.DB) {
	tableStatements := []string{
		`CREATE TABLE IF NOT EXISTS users (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
        	username VARCHAR(255),
        	email VARCHAR(255),
            password VARCHAR(255),
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
		`CREATE TABLE IF NOT EXISTS pages (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title VARCHAR(255),
			user_id INTEGER,
            parent_id INTEGER,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (parent_id) REFERENCES pages (id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
        )`,
		`CREATE INDEX IF NOT EXISTS parent_id_index ON pages (parent_id)`,
		`CREATE TABLE IF NOT EXISTS blocks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            page_id INTEGER NOT NULL,
            type VARCHAR(255) NOT NULL,
            content TEXT,
            position INTEGER NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (page_id) REFERENCES pages (id) ON DELETE CASCADE
        )`,
		`CREATE TABLE IF NOT EXISTS page_properties (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            page_id INTEGER NOT NULL,
            name VARCHAR(255) NOT NULL,
            value TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (page_id) REFERENCES pages (id) ON DELETE CASCADE
        )`,
		`CREATE TABLE IF NOT EXISTS page_templates (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name VARCHAR(255) NOT NULL,
            structure TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
		// `INSERT INTO users (username, email, password) VALUES ("duyhung", "duyhung@gmail.com", "duyhung")`,
		// `INSERT INTO pages (user_id, title) VALUES (1, "First page")`,
	}

	ctx := context.Background()
	for _, stmt := range tableStatements {
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
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Username.GetOr(""))
	}
}
