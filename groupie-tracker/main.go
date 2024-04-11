package main

import (
	"groupie-tracker/data"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", handlergaylish)
	http.ListenAndServe(":8080", nil)
}

func handlergaylish(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	artist, _ := data.Responce()
	t.Execute(w, artist)
}
