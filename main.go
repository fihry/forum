package main

import (
	"log"
	"net/http"

	"forum/api/controllers"
	"forum/api/routes"
	"forum/utils"
)

func main() {
	port, err := utils.GetPort()
	if err != nil {
		log.Println(err)
	}
	// create a new server
	server := http.Server{
		Addr:    "0.0.0.0" + port,
		Handler: routes.InitRouter(),
	}

	// print the location of the server
	log.Println("\033[32mServer is running on http://localhost" + port + "...🚀\033[0m")

	// init database
	err = controllers.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	// defer close the database connection
	defer controllers.Database.DB.Close()

	// start the server
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
