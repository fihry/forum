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
	log.Println("Server started at port", Port)
	forum.InitDB()
	err = http.ListenAndServe(Port, nil)
	if err != nil {
		panic(err)
	}
}
