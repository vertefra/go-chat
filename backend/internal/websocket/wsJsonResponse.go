package websocket

import "log"

// Define the response from a websocket
type WsJsonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

func broadcastToAll(response WsJsonResponse) {
	for client := range Clients {
		err := client.WriteJSON(response)

		if err != nil {
			log.Println("Websocket error")
			_ = client.Close()
			delete(Clients, client)
		}
	}
}
