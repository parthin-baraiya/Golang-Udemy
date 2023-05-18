package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	movieId := mux.Vars(req)["id"]

	for _, movie := range movies {
		if movie.ID == movieId {
			json.NewEncoder(res).Encode(movie)
			break
		}
	}
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	movieId := mux.Vars(req)["id"]

	for i, movie := range movies {
		if movie.ID == movieId {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

	json.NewEncoder(res).Encode(movies)

}

func createMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(req.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, newMovie)
	json.NewEncoder(res).Encode(newMovie)
}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	movieId := mux.Vars(req)["id"]

	var newMovie Movie
	_ = json.NewDecoder(req.Body).Decode(&newMovie)

	for i, movie := range movies {
		if movie.ID == movieId {
			movies[i].Title = newMovie.Title
			movies[i].Director = newMovie.Director
			break
		}
	}

	json.NewEncoder(res).Encode(movies)
}

func main() {

	movies = append(movies, Movie{ID: "1", Title: "Hello world", Director: &Director{FirstName: "Parthin", LastName: "Baraiya"}})
	movies = append(movies, Movie{ID: "2", Title: "Apple", Director: &Director{FirstName: "steve", LastName: "job"}})

	server := mux.NewRouter()

	server.HandleFunc("/movies", getMovies).Methods("GET")
	server.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	server.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	server.HandleFunc("/movies", createMovie).Methods("POST")
	server.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	fmt.Printf("Server is running on port 8000\n")
	if err := http.ListenAndServe(":8000", server); err != nil {
		log.Fatal(err)
	}
}
