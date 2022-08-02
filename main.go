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

type (
	Movie struct {
		ID       string    `json:"id"`
		Isbn     string    `json:"isbn"`
		Title    string    `json:"title"`
		Director *Director `json:"director"`
	}
	Director struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}
)

var movies []Movie

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Conternt-Type", "application/json")
	parameters := mux.Vars(request)
	for index, value := range movies {
		if value.ID == parameters["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(movies)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request)
	for _, value := range movies {
		if value.ID == parameters["id"] {
			json.NewEncoder(writer).Encode(value)
			return
		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Conternt-Type", "application/json")
	parameters := mux.Vars(request)
	for index, value := range movies {
		var movie Movie
		if value.ID == parameters["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = parameters["id"]
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movie)
			return
		}
	}

}

func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "48156", Title: "The Hunt", Director: &Director{Firstname: "Thomas", Lastname: "Vinterberg"}})
	movies = append(movies, Movie{ID: "2", Isbn: "481567", Title: "Another Round", Director: &Director{Firstname: "Thomas", Lastname: "Vinterberg"}})
	route := mux.NewRouter()
	route.HandleFunc("/movies", getMovies).Methods("GET")
	route.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	route.HandleFunc("/movies", createMovie).Methods("POST")
	route.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	route.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 9000\n")
	log.Fatal(http.ListenAndServe(":9000", route))
}
