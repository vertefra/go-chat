package web

import (
	"log"
	"net/http"
)

func Web() {
	router := routes()

	log.Println("Starting server on port 8080")

	http.ListenAndServe(":8080", router)
}
