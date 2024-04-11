package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	LocationsData
}

type LocationsData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

const (
	ArtistUrl    string = "https://groupietrackers.herokuapp.com/api/artists"
	RelationsUrl string = "https://groupietrackers.herokuapp.com/api/relation/"
)

func Responce() ([]Artists, error) {
	var artists []Artists
	zapros, err := http.Get(ArtistUrl)
	if err != nil {
		fmt.Println("Ne schital json", err)
		return nil, err
	}
	defer zapros.Body.Close()

	if zapros.StatusCode != http.StatusOK {
		return nil, err
	}

	err = json.NewDecoder(zapros.Body).Decode(&artists)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range artists {
		v.LocationsData, err = ResponceRelations(v.Id)
		if err != nil {
			return nil, err
		}
	}

	// for _, ch := range artists {
	// 	fmt.Println(ch.Name)
	// }
	//	fmt.Println(artists)
	return artists, nil
}

func ResponceRelations(id int) (LocationsData, error) {
	abob := strconv.Itoa(id)
	var relation LocationsData
	res, err := http.Get(RelationsUrl + abob)
	if err != nil {
		fmt.Println("Ошибка считывания Json")
		return relation, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&relation)
	if err != nil {
		fmt.Println("Ошибка декордирование", err)
		return relation, err
	}
	// for _, ch := range relation {
	// 	fmt.Println(ch.ID)
	// }
	return relation, nil
}
