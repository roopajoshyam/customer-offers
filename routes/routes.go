package routes

import (
	"github.com/gorilla/mux"
	"github.com/jroopa/offers/controllers"
	"github.com/jroopa/offers/middlewares"
)


func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(middlewares.CommonMiddleware)
	r.HandleFunc("/", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/create/offer", controllers.CreateOffer).Methods("POST")

	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middlewares.JwtVerify)
	s.HandleFunc("/user", controllers.FetchUsers).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	s.HandleFunc("/users/offers", controllers.GetUserOffers).Methods("GET")

	return r
}
