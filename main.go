package main

import (
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
	Mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1" + Port,
		Handler: Mux,
	}
	// serve the css files
	Mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/css"))))
	// handle routes
	Mux.HandleFunc("/", forum.HomePageHandler)
	Mux.HandleFunc("/login", forum.LoginPageHandler)
	Mux.HandleFunc("/register", forum.RegisterPageHandler)

	// print the location of the server
	log.Println("\033[32mServer is running on port " + Port + "...🚀\033[0m")
	log.Println("\033[32mhttp://localhost" + Port + "\033[0m")
	// init database
	forum.InitDB()
	// start the server
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
