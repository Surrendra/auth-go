package main

import (
	"log"
	"net/http"

	"github.com/Surrendra/auth-go/controllers/authcontroller"
	"github.com/Surrendra/auth-go/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase()
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
