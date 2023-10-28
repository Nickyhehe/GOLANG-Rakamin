package router

import (
	"github.com/gorilla/mux"
	"github.com/username/projectname/controllers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/login", controllers.LoginUser).Methods("POST")

	router.HandleFunc("/photos/{id}", controllers.GetPhoto).Methods("GET")
	router.HandleFunc("/photos/{id}", controllers.DeletePhoto).Methods("DELETE")
	router.HandleFunc("/photos/{id}", controllers.UpdatePhoto).Methods("PUT")

	return router
}
