package main

import (
	"log"
	"net/http"
)

var server *http.Server = &http.Server{
	Addr:    "localhost:8000",
	Handler: nil,
}

func main() {
	log.Printf("Starting Server at http://%s\n", server.Addr)

	if error := server.ListenAndServe(); error != nil {
		log.Fatal("Something Went Wrong", error)
	}
}
