package web

import (
	"log"
	"net/http"
	ws "vertefra/go-chat/internal/websocket"
)

func Web() {
	router := routes()

	log.Println("Starting channel listener")

	go ws.ListenToWSChannel()

	log.Println("Starting server on port 8080")

	http.ListenAndServe(":8080", router)
}
