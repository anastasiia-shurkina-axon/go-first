package main

import (
	"log"
	"net/http"

	"github.com/anastasiia-shurkina-axon/go-first/handlers"
)

func main() {
	log.Print("Starting REST API service")
	s := handlers.NewServer()
	log.Print("Listening on port :8081")
	log.Fatal(http.ListenAndServe(":8081", s.GetRouter()))
}
