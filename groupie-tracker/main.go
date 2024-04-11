package main

import (
	"fmt"
	"groupie-tracker/data"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handlergaylish)
	http.HandleFunc("/artist", artistHandler)
	http.ListenAndServe(":8080", nil)
}

func handlergaylish(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	artists, _ := data.Responce()
	t.Execute(w, artists)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	artists, _ := data.Responce()

	artist, err := aboba(artists, id)
	if err != nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	t, _ := template.ParseFiles("artist.html")

	t.Execute(w, artist)
}

func aboba(artists []data.Artists, id int) (data.Artists, error) {
	var artist data.Artists
	fmt.Println(id)
	for _, ch := range artists {
		fmt.Println(ch.Id)
		if ch.Id == id {
			artist = ch
		}
	}
	return artist, nil
}
