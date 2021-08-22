package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var wsChan = make(chan WsPayload)

var clients = make(map[WebSocketConnection]string)

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

	// just wrapping ws with WebsocketConnection struct
	conn := WebSocketConnection{Conn: ws}

	// setting the map using conn as a key
	clients[conn] = ""

	err = ws.WriteJSON(response)

	if err != nil {
		log.Println(err)
	}

	go ListenForWS(&conn)
}

func ListenForWS(conn *WebSocketConnection) {
	// If this function will panic inside the go runtine
	// this deferred function will be executed to recover
	// ListenForWS
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	// In this infinite loop we keep checking if
	// Payload arrives on conn it sends it on the channel
	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func ListenToWSChannel() {
	var response WsJsonResponse

	for {
		e := <-wsChan
		response.Action = "Got here"
		response.Message = fmt.Sprintf("Some message %s", e.Action)
	}
}

func broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)

		if err != nil {
			log.Println("Websocket error")
			_ = client.Close()
			delete(clients, client)
		}
	}
}

type WebSocketConnection struct {
	*websocket.Conn
}

// Define the response from a websocket
type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type WsPayload struct {
	Action   string              `json:"action"`
	Message  string              `json:"message"`
	Username string              `json:"username"`
	Conn     WebSocketConnection `json:"-"`
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
