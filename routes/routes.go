package routes

import (
	"GolangAPIPractice/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	// router
	subRouter := router.PathPrefix("/users").Subrouter()
	subRouter.HandleFunc("/all", controllers.ListUsers).Methods("GET")
	subRouter.HandleFunc("/{id}", controllers.GetUserDetail).Methods("GET")
	subRouter.HandleFunc("/createUser", controllers.CreateUser).Methods("POST")
	// Server static file with prefix static
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	return router
}
