package main

import (
	"log"
	"net/http"
)

func main() {
	router := routes()

	log.Println("Starting server on port 8080")

	_ = http.ListenAndServe(":8080", router)
}
