package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kawaljeet2001/netflix-api/helpers"
	"github.com/Kawaljeet2001/netflix-api/model"
	"github.com/gorilla/mux"
)

func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := helpers.GetAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var createdMovie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&createdMovie)

	if err != nil {
		log.Fatal(err)
	}
	helpers.CreateMovie(createdMovie)
	json.NewEncoder(w).Encode(createdMovie)
}

func SetMovieAsWatchedController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	helpers.UpdateMovieById(params["movieId"])

	json.NewEncoder(w).Encode(params["movieId"])
}

func DeleteMovieByIdController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	helpers.DeleteMovieById(params["movieId"])

	responseObject := make(map[string]string)

	responseObject["message"] = "Deleted the movie successfully"
	responseObject["deletedId"] = params["movieId"]
	json.NewEncoder(w).Encode(responseObject)
}

func DeleteAllMoviesEncoder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	helpers.DeleteAllMovies()

	responseObject := make(map[string]string)

	responseObject["message"] = "Deleted all movies successfully"
	json.NewEncoder(w).Encode(responseObject)

}
