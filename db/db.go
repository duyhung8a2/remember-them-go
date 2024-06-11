package db

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/stephenafamo/bob"
	_ "modernc.org/sqlite"
)

const (
	DBDriver     = "sqlite"
	DBFile       = "db/database.db"
	DBSchemaFile = "db/db_schema.sql"
)

func ConnectDB() (*bob.DB, error) {
	db, err := bob.Open(DBDriver, DBFile)
	if err != nil {
		return nil, err
	}
	return &db, nil
}

func InitDB(db *bob.DB) {
	content, err := os.ReadFile(DBSchemaFile)
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
