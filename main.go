package main

import (
	"log"
	"net/http"

	"github.com/jalopezma/shops-api/handlers"
)

func main() {
	router := handlers.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
