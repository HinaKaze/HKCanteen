package handler

import (
	"html/template"
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./client/index.html")
	log.Println("welcome invoke")
	t.Execute(w, nil)
}
