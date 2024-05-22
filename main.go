package main

import (
	"github.com/google/uuid"
	"log"
	"net/http"

	"github.com/cheeseNA/owlback/config"
	ogen "github.com/cheeseNA/owlback/ogen"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	log.Println(cfg.PostgresConnectionString)

	service := &service{
		tasks: map[uuid.UUID]ogen.Task{},
	}
	// Create generated server.
	srv, err := ogen.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
