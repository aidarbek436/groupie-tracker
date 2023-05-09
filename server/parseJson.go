package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Parse_Json() *[]Artist {
	artistPage, err := http.Get("https://groupietrackers.herokuapp.com/api/artists") // fetching URL
	if err != nil {
		log.Fatal(err)
	}
	defer artistPage.Body.Close()                   // must have
	content, err := ioutil.ReadAll(artistPage.Body) // reading body and saving
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &Artists) // sending json data into struct
	if err != nil {
		log.Fatal(err)
	}
	return &Artists
}

func parseLocationsJson_AddToStruct(id_string string, id_int int) {
	read_API, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id_string) // fetching URL
	if err != nil {
		log.Fatal("Fetch issue")
	}
	defer read_API.Body.Close()                     // must have
	read_Json, err := ioutil.ReadAll(read_API.Body) // reading body and saving
	if err != nil {
		log.Fatal("Reading json issue")
	}
	var data Relations
	err = json.Unmarshal(read_Json, &data) // sending json data into created struct with map
	if err != nil {
		log.Fatal("Unmarshal json issue")
	}
	Artists[id_int-1].DatesLocations = data.DatesLocations // saving map data into main struct
}
