package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	Id       string    `json"id"`
	Isbn     string    `json"isbn"`
	Title    string    `json"title"`
	Director *Director `json"director"`
}

type Director struct {
	Firstman string `json:"firstman"`
	Lastman  string `json:"lastman"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for index, item := range movies {
		if item.Id == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, item := range movies {
		if item.Id == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func creatMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//phram
	param := mux.Vars(r)

	for index, item := range movies {
		if item.Id == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = param["id"]
			movies = append(movies, movie)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "2001", Title: "Movie One", Director: &Director{Firstman: "lightning", Lastman: "McQueen"}})
	movies = append(movies, Movie{Id: "2", Isbn: "2002", Title: "Movie two", Director: &Director{Firstman: "Bruce", Lastman: "Wayne"}})
	r.HandleFunc("movies", getMovies).Methods("GET")
	r.HandleFunc("movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", creatMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
