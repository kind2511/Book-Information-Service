package main

import (
	"assignment-1/handlers"
	"assignment-1/utilities"
	"log"
	"net/http"
	"os"
)

func main() {

	// Assign Port
	port:= os.Getenv("PORT");

	if port == "" {
		log.Println("The PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Handler Endpoints
	http.HandleFunc(utilities.DEFAULT_PATH, handlers.EmptyHandler)


	// Start http Server
	log.Println("Starting server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

