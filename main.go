package main

import (
	"log"
	"net/http"

	"github.com/Surrendra/auth-go/controllers/authcontroller"
	"github.com/Surrendra/auth-go/controllers/productcontroller"
	"github.com/Surrendra/auth-go/middlewares"

	"github.com/Surrendra/auth-go/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase()
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/product", productcontroller.Index).Methods("GET")
	api.HandleFunc("/product/{id}", productcontroller.Find).Methods("GET")
	api.HandleFunc("/product/create", productcontroller.Create).Methods("POST")
	api.HandleFunc("/product/update/{id}", productcontroller.Update).Methods("PUT")
	api.HandleFunc("/product/delete/{id}", productcontroller.Delete).Methods("POST")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
