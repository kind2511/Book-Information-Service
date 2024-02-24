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
	http.HandleFunc(utilities.DEFAULT_PATH, handlers.DefaultHandler)
	http.HandleFunc(utilities.STATUS_PATH, handlers.StatusHandler)
	http.HandleFunc(utilities.BOOKCOUNT_PATH, handlers.BookCountHanlder)
	http.HandleFunc(utilities.READERSHIP_PATH, handlers.ReadershipHandler)

	// Start http Server
	log.Println("Starting server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}