# Go chat

Simple example of real time chat using gorilla websocket package with go.

Client request the connection with server creating the socket 

```javascript
socket = new WebSocket("ws://127.0.0.1:8080/ws")
```

When receiving the request on `/ws` endpoint, the server upgrades the connection through the `Upgrade` method defined on the `Upgrader` struct,
saves the client connection as key in a map that will eventually holds the username as a value, and launch a goroutine that will constantly check if any payload arrives on this connection (websocket.ListenForWS). If a payload is received will be sent to the `wsChan` channel

Another goroutine, launched in web main.go, ListenToWSChannel, constantly monitors if any payload is received through `wsChan` channel. 

When a payload is received, based on the `Action` case, a response is sent. In this case only a `BroadcastToAll` method has been implemented, that will send the response to all the clients in the Client map.