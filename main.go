package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

type song struct {
	ID     string `json:"ID"`
	Title  string `json:"Title"`
	Artist string `json:"Artist"`
}

var library []song

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called homePage()")
}

func getLibrary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint called getLibrary()")

	json.NewEncoder(w).Encode(library)
}

func addSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newSong song
	_ = json.NewDecoder(r.Body).Decode(&newSong)

	library = append(library, newSong)

	json.NewEncoder(w).Encode(newSong)
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := mux.Vars(r)

	_deleteID(p["id"])

	json.NewEncoder(w).Encode(library)
}

func _deleteID(id string) {
	for index, song := range library {
		if song.ID == id {
			library = append(library[:index], library[index+1:]...)
			break
		}
	}
}

func updateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Song song
	_ = json.NewDecoder(r.Body).Decode(&Song)

	p := mux.Vars(r)

	_deleteID(p["id"])
	library = append(library, Song)

	json.NewEncoder(w).Encode(library)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/library", getLibrary).Methods("GET")
	router.HandleFunc("/library", addSong).Methods("POST")
	router.HandleFunc("/library/{id}", deleteSong).Methods("DELETE")
	router.HandleFunc("/library/{id}", updateSong).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	library = append(library, song{
		ID:     "0",
		Title:  "Chant Down Babylon",
		Artist: "Bob Marrley",
	})

	handleRequests()
}
