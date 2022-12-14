package main

import (
	"github.com/anastasiia-shurkina-axon/go-first/file/handlers"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting File REST API service")
	s := handlers.NewServer()
	log.Print("Listening on port :8081")
	log.Fatal(http.ListenAndServe(":8081", s.GetRouter()))
}
