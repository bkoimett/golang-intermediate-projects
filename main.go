package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbm"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/jsom")
	json.NewEncoder(w).Encode(movies)
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "43277", Title: "Dr.Doomsday", Director: &Director{Firstname: "Eddy", Lastname: "Murphy"}})
	movies = append(movies, Movie{ID: "2", Isbn: "24199", Title: "Family Reunion", Director: &Director{Firstname: "Tyler", Lastname: "Perry"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("server has started on https://localhost:8000 ...")
	log.Fatal(http.ListenAndServe(":8000", r))

}