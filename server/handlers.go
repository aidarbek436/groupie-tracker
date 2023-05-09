package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Artist struct {
	Id             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	Relations      string   `json:"relations"`
	DatesLocations map[string][]string
}

var Artists []Artist

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404:Page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "405:Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	Artists := Parse_Json()
	template, err := template.ParseFiles("templates/mainpage.html")
	if err != nil {
		fmt.Println("template parse fail")
		return
	}
	err = template.Execute(w, Artists)
	if err != nil {
		fmt.Println("template execute fail")
		return
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	artist_id_string := r.URL.Query().Get("id")
	artist_id_int, err := strconv.Atoi(artist_id_string)
	if err != nil {
		http.Error(w, "404:Page not found", http.StatusNotFound)
		return
	}
	if artist_id_int < 1 || artist_id_int > 52 {
		http.Error(w, "404:Page not found", http.StatusNotFound)
		return
	}
	parseLocationsJson_AddToStruct(artist_id_string, artist_id_int)
	template, err := template.ParseFiles("templates/artistpage.html")
	if err != nil {
		fmt.Println("template parse fail")
		return
	}
	err = template.Execute(w, Artists[artist_id_int-1])
	if err != nil {
		fmt.Println("template execute fail")
		return
	}
}
