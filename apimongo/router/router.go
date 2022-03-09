package router

import (
	"github.com/gorilla/mux"
	"github.com/jeeshan12/apimongo/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", controller.ServeHome).Methods("GET")
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/deleteAllMovies", controller.DeleteAllMovies).Methods("DELETE")
	return router
}
