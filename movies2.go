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
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	//set headers
	w.Header().Set("Content-Type:", "application/json")

	//error validation
	if (r.Method != "GET"){
		fmt.Printf("invalid request method")
	}

	//error validation
	if (r.URL.Path != "/movies"){
		fmt.Printf("invalid request method")
	}

	//get all movies
	json.NewEncoder(w).Encode(movies)
	fmt.Fprintf(w, "Movies: %v", movies)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	//set headers
	w.Header().Set("Content-Type", "application/json")

	//create movie variable
	var movie Movie

	//create new movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)

	//return movie
	json.NewEncoder(w).Encode(movie)

	// fmt.Fprintf(w, "Created Movie: %v", movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	//set headers
	w.Header().Set("Content-Type", "application/json")

	//create movie variable
	var movie Movie

	//get params from vars
	params := mux.Vars(r)

	//cycle through movies slice and delete specified id
	for index, item := range movies {
		if (item.ID == params["id"]){
			return movies(movies[:index], movies[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(movie)

}

func getMovie(w http.ResponseWriter, r *http.Request){
	//set headers 
	w.Header().Set("Content-Type", "application/json")

	//get the params from vars
	params := mux.Vars(r)

	//create movie variable
	var movie Movie

	//cycle through movies slice and compare params id to movie id
	for index, item := range movies {
		if (item.ID == params["id"]){
			return index[movie]
		}
	}
	json.NewEncoder(w).Encode(movie)
}


func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "2", Title: "Avengers", Director: &Director{Firstname: "Johnson", Lastname: "Brothers"}})
	movies = append(movies, Movie{ID: "2", Isbn: "3", Title: "Avatar", Director: &Director{Firstname: "James", Lastname: "Cameron"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies", updateMovie).Methods("PATCH")
	// r.HandleFunc("/movies", deleteMovie).Methods("DELETE")

	fmt.Printf("server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r)) 

}