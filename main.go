package main

import (
	"forum/api/Routes"
	forum "forum/src"
	"log"
	"net/http"
)

func main() {
	Port, err := forum.GetPort()
	if err != nil {
		log.Println(err)
	}
	// create a new server
	server := http.Server{
		Addr:    "127.0.0.1" + Port,
		Handler: forum.Routs(),
	}
	// print the location of the server
	log.Println("\033[32mServer is running on port " + Port + "...🚀\033[0m")
	log.Println("\033[32mhttp://localhost" + Port + "\033[0m")
	// init database
	db, err := forum.InitDB()
	if err != nil {
		log.Println(err)
	}
	// create a new database
	Routes.Database.DB = db
	// close the database connection
	defer db.Close()
	// handle auth routes
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
