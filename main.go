package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	fmt.Printf("Server is running on port 8000\n")
	log.Fatal(http.ListenAndServe(":8002", r))
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	moviesList := movies // This retrieves the list of movies

	// Prepare the response structure
	response := map[string]interface{}{
		"success": true,
		"data":    moviesList, // This includes the movie list under 'data'
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode and send the JSON response
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
