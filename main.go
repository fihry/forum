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

	log.Println("\033[32mServer running...\033[0m")
	log.Println("\033[32mThe server has been launched at: " + server.Addr + "/\033[0m")
	// init database
	forum.InitDB()

	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
