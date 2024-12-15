package utils

import (
	"html/template"
	"log"
	"net/http"
)


func RenderTemplate(w http.ResponseWriter, file string, data interface{}) {
	t, err := template.ParseFiles("./view/template/" + file + ".html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}
