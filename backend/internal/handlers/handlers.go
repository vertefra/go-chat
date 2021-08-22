package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// Upgrade the connection
var upgradeConnections = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Generic Websocket endpoint with upgraded connection
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnections.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected to WS endpoint")

	var response WsJsonResponse
	response.Message = `<em><small>Connected to server</small></em>`
	err = ws.WriteJSON(response)

	if err != nil {
		log.Println(err)
	}
}

// Define the response from a websocket
type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
