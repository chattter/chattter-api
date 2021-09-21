package services

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// SocketMessage is a message sent/received across the socket
type SocketMessage struct {
	// Type is the type of action being performed by this message
	Type string `json:"type"`
	// Sha is the SHA256 hash of the JSON contents of the rest of the message
	Sha string `json:"sha"`
	// ReplySha is the SHA256 hash of the message that is being replied to
	ReplySha *string `json:"reply_sha"`
	// Data is the data for the message itself, formatted according to the `Type`
	Data interface{} `json:"data"`
}

// SocketSession contains all of the session-specific details for the socket conection, such as authentication,
// rooms, etc
type SocketSession struct {
	// ...
}

// SocketConn represents a single connected client
type SocketConn struct {
	// key is a unique key that is assigned by the server and used to identify the client connection. It is never
	// shared with the frontend client, as it has no utility to the client
	key uint64
	// conn is the raw WebSocket connection initiated by the client
	conn *websocket.Conn
	// server is the SocketsService with which this client is registered. Useful if we need to broadcast messages
	// to other connected clients
	server *SocketsService
	// session contains all of the data specific to this session (authentication, rooms, etc.)
	session SocketSession
}

// Run loops indefinitely, handling messages sent by the client, until the client disconnects
func (c *SocketConn) Run() error {

	// Loop indefinitely reading messages from the client
	for {

		// Read a message from the client
		var msg SocketMessage
		if err := c.conn.ReadJSON(&msg); err != nil {

			// If it's a normal closure of the socket by the client, exit the loop
			// and return without error
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				break
			}

			// Otherwise, return the error
			// TODO: Recover from some errors, so the client doesn't get disconnected
			// for minor things
			return err

		}

		// Handle the message, and log any errors
		if err := c.HandleMessage(&msg); err != nil {
			fmt.Println("Error handling socket message: ", err.Error())
		}

	}

	// Return without error
	return nil

}

// HandleMessage interprets a single message received from the client, and responds to it
func (c *SocketConn) HandleMessage(msg *SocketMessage) error {

	// The message type is unknown
	return fmt.Errorf("unrecognized message type: %s", msg.Type)

}
