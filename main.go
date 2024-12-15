package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"forum/api/controllers"
	"forum/api/routes"
	"forum/utils"
)

func main() {
	// set the usage flag for the command line -h or --help
	flag.Usage = func() {
		fmt.Println("Usage: forum [flags]")
		flag.PrintDefaults()
	}
	// set the log prefix
	log.SetPrefix("\033[34m[forum]-[\033[0m")
	// load the config
	config := utils.LoadConfig()
	// set the address of the server
	Adress := config.Server.Host + ":" + config.Server.DefaultPort
	// create a new server
	server := http.Server{
		Addr:    Adress,
		Handler: routes.InitRouter(),
	}

	// print the location of the server
	log.Printf("\033[32m] Server is running...🚀\nLink: 🌐 http://%s\033[0m\n", server.Addr)

	// init database
	err := controllers.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	// defer close the database connection
	defer controllers.Database.Close()

	// start the server
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
