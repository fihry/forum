package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDB() error {
	var err error
	Database, err = sql.Open("sqlite3", "db/Forum.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	err = Database.Ping()
	if err != nil {
		log.Panic()
	}

	// read the db schema from the schema.sql file
	schema, err := os.ReadFile("db/migrations/schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}
	// execute the schema
	_, err = Database.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}
	return nil
}
