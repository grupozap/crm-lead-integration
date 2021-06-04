package main

import (
	"fmt"
	"log"
	"net/http"

	"leadIntegration/src/auth"
	"leadIntegration/src/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	router := mux.NewRouter()

	router.Use(auth.BasicAuthentication)

	router.HandleFunc("/health-check", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/leads/lead", controllers.RecieveLead).Methods("POST")

	log.Print("Server up on port: 8000")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}
}
