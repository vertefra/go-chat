package web

import (
	"log"
	"net/http"
	"vertefra/go-chat/internal/handlers"
)

func Web() {
	router := routes()

	log.Println("Starting channel listener")

	go handlers.ListenToWSChannel()

	log.Println("Starting server on port 8080")

	http.ListenAndServe(":8080", router)
}
