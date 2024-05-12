package main

import (
	"log"
	"net/http"
	"stakeholders-service/controller"
	"stakeholders-service/service"

	"github.com/gorilla/mux"
)

func startServer(tokenController *controller.JWTController) {
	router := mux.NewRouter().StrictSlash(true)

	//TOURS
	//router.HandleFunc("/tours/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/token", tokenController.GenerateAccsessToken).Methods("POST")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", router))
		
}

func main() {


	//equipment
	
	jwtService := &service.JWTService{}
	jwtController := &controller.JWTController{JwtService: jwtService}

	

	startServer(jwtController)
}

