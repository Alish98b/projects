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

type LocationsData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

const (
	ArtistUrl    string = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsUrl string = "https://groupietrackers.herokuapp.com/api/locations"
)

func Responce() ([]Artists, error) {
	var artists []Artists
	zapros, err := http.Get(ArtistUrl)
	if err != nil {
		fmt.Println("Ne schital json", err)
		return nil, err
	}
	defer zapros.Body.Close()

	decoder := json.NewDecoder(zapros.Body)

	err = decoder.Decode(&artists)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// for _, ch := range artists {
	// 	fmt.Println(ch.Relations)
	// }
	return artists, nil
}

func ResponceRelations() ([]LocationsData, error) {
	var locations []LocationsData
	res, err := http.Get(LocationsUrl)
	if err != nil {
		fmt.Println("Ошибка считывания Json")
		return nil, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&locations)
	if err != nil {
		fmt.Println("Ошибка декордирование", err)
		return nil, err
	}
	for _, ch := range locations {
		fmt.Println(ch.ID)
	}
	return locations, nil
}
