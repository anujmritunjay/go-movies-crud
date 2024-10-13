package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}
type Response struct {
	Success bool    `json:"success"`
	Data    []Movie `json:"data"`
}

var movies []Movie

func main() {

	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Title for movie1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45678", Title: "Title for movie2", Director: &Director{Firstname: "Steven", Lastname: "Smith"}})
	r := mux.NewRouter()
	r.HandleFunc("/", rootRoute).Methods("GET")
	r.HandleFunc("/get-movies", getAllMovies).Methods("GET")
	r.HandleFunc("/get-movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/add-movie", addMovie).Methods("POST")
	r.HandleFunc("/delete-move/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/update-movie/{id}", updateMovie).Methods("PUT")

	fmt.Printf("Server is running on port 8000\n")
	log.Fatal(http.ListenAndServe(":8002", r))
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movieId := mux.Vars(r)
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	for index, val := range movies {
		if val.ID == movieId["id"] {
			movies[index].Director = movie.Director
			movies[index].Title = movie.Title
			movies[index].Director = movie.Director
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"data":    movies[index],
		})
		return
	}

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	for index, movie := range movies {
		if movie.ID == movieId {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Movie deleted successfully."})
		}
	}

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	for _, movie := range movies {
		if movie.ID == movieId {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.IntN(1000000))
	movies = append(movies, movie)

	res := Response{
		Success: true,
		Data:    movies,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	moviesList := movies
	response := map[string]interface{}{
		"success": true,
		"data":    moviesList,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func rootRoute(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"success": true,
		"message": "Hello from the Go lang server",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
