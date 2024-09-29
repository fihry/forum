package main

import (
	forum "forum/src"
	"net/http"
)

func main() {
	Port := ":8080"
	forum.InitDB()
	http.ListenAndServe(Port, nil)
}
