package main

import (
	"log"
	"net/http"
	"student/group/server"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.MainPage)
	mux.HandleFunc("/artist/", server.ArtistPage)
	log.Println("server : http:localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
