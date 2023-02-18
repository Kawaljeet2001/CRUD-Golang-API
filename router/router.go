package router

import (
	"github.com/Kawaljeet2001/netflix-api/controller"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServeHomeController).Methods("GET")
	router.HandleFunc("/movies", controller.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/movie", controller.CreateMovieController).Methods("POST")
	router.HandleFunc("/movie/{movieId}", controller.SetMovieAsWatchedController).Methods("PUT")
	router.HandleFunc("/movie/{movieId}", controller.DeleteMovieByIdController).Methods("DELETE")
	router.HandleFunc("/moviesd", controller.DeleteAllMoviesController).Methods("DELETE")

	return router
}
