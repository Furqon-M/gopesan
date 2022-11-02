package main

import (
	"log"
	"net/http"

	"github.com/Furqon-M/gopesan/controllers/authcontroller"
	"github.com/Furqon-M/gopesan/controllers/menucontroller"
	"github.com/Furqon-M/gopesan/controllers/pesancontroller"
	"github.com/Furqon-M/gopesan/middlewares"
	"github.com/Furqon-M/gopesan/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/menu", menucontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	r.HandleFunc("/menu/{id}", menucontroller.Show).Methods("GET")
	r.HandleFunc("/menu", menucontroller.Create).Methods("POST")
	r.HandleFunc("/menu/{id}", menucontroller.Update).Methods("PUT")
	r.HandleFunc("/menu", menucontroller.Delete).Methods("DELETE")

	r.HandleFunc("/pesan", pesancontroller.Index).Methods("GET")
	r.HandleFunc("/pesan/{id}", pesancontroller.Show).Methods("GET")
	r.HandleFunc("/pesan", pesancontroller.Create).Methods("POST")
	r.HandleFunc("/pesan/{id}", pesancontroller.Update).Methods("PUT")
	r.HandleFunc("/pesan", pesancontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
