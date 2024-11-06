package controllers

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDB() error {
	db, err := sql.Open("sqlite3", "db/Forum.db")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// read the database schema from the schema.sql file
	schema, err := os.ReadFile("db/schema.sql")
	if err != nil {
		return err
	}
	// execute the schema
	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	// set the database to the database object
	Database = db
    log.Println("✅ Database initialized")
	return nil
}
