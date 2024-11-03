package controllers

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDB() error {
	db, err := sql.Open("sqlite3", "./db/Forum.db")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return errors.New("error in  db connection:  \n" + err.Error())
	}

	log.Println("\033[32mConnected to database successfully\033[0m")

	// Create tables if not exists
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(20) UNIQUE NOT NULL,
        password TEXT NOT NULL,
        email TEXT NOT NULL,
        session TEXT NOT NULL
    )
`)
	if err != nil {
		return errors.New("error in  creating table users: \n" + err.Error())
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
        author INTEGER NOT NULL,
        category TEXT NOT NULL,
        likes INTEGER,
        dislikes INTEGER,
        Foreign key (author) references users(id)
    )
`)
	if err != nil {
		return errors.New("error in creating table posts: \n" + err.Error())
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY,
        postId INTEGER,
        author INTEGER NOT NULL,
        content TEXT,
        FOREIGN KEY (postId) REFERENCES posts(id),
        FOREIGN KEY (author) REFERENCES users(id)
    )
`)
	if err != nil {
		return errors.New("error in creating table comments:  \n" + err.Error())
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS engagement (
    id INTEGER PRIMARY KEY,
    postId INTEGER,
    userId INTEGER,
    commentId INTEGER,
    like BOOLEAN,
    dislike BOOLEAN,
    FOREIGN KEY (postId) REFERENCES posts(id),
    FOREIGN KEY (userId) REFERENCES users(id),
    FOREIGN KEY (commentId) REFERENCES comments(id)
    )
`)
	if err != nil {
		return errors.New("error in creating  table engagement: \n" + err.Error())
	}

	// set the database to the database object
	Database = db

	return nil
}
