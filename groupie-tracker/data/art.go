package data

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

const ArtistURL string = "https://groupietrackers.herokuapp.com/api/artists"

func Responce() ([]Artists, error) {
	var artists []Artists
	zapros, err := http.Get(ArtistURL)
	if err != nil {
		fmt.Println("Ne schital json")
		return nil, err
	}
	defer zapros.Body.Close()

	decoder := json.NewDecoder(zapros.Body)

	err = decoder.Decode(&artists)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, ch := range artists {
		fmt.Println(ch.Name)
	}
	return artists, nil
}
