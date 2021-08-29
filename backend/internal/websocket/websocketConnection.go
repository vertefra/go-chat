package websocket

import (
	"fmt"
	"log"
	"sort"

	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	*websocket.Conn
}

var wsChan = make(chan WsPayload)
var Clients = make(map[WebSocketConnection]string)

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
		switch e.Action {

		case "left":
			response.Action = "list_users"
			delete(Clients, e.Conn)
			users := getUserLists()
			response.ConnectedUsers = users
			broadcastToAll(response)

		case "username":
			// registers the username in the active clients
			Clients[e.Conn] = e.Username
			users := getUserLists()
			response.Action = "list_users"
			response.ConnectedUsers = users
			broadcastToAll(response)

		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf(
				"<strong>%s</strong>: %s \n", e.Username, e.Message,
			)

			broadcastToAll(response)
		}

	}
}

func getUserLists() []string {
	var userList []string

	for _, x := range Clients {
		if x != "" {
			userList = append(userList, x)
		}
	}

	sort.Strings((userList))

	return userList
}
