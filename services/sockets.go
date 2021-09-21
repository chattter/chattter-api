package services

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// upgrader is the websocket upgrader instance with default configuration
var upgrader websocket.Upgrader

// randKeySource is the random source used to generate random values
var randKeySource = rand.NewSource(time.Now().Unix())

// SocketsService provides a WebSockets server and manages connected clients
type SocketsService struct {
	// clients is a map of clients actively connected to the server
	clients map[uint64]*SocketConn
	// clientsMut is a mutex to control concurrent access to the connected clients
	clientsMut sync.RWMutex
}

// ServeHTTP handles upgrading an inbound WebSocket request from HTTP to WS protocol
func (s *SocketsService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// Attempt to upgrade the caller
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer conn.Close()

	// Create the wrapped socket connection and add it to the map of all
	// active connections
	client := SocketConn{
		conn:   conn,
		server: s,
	}
	if err := s.addClient(&client); err != nil {
		fmt.Println("Failed to add client: ", err.Error())
		return
	}

	// Run the client indefinitely
	if err := client.Run(); err != nil {
		fmt.Println("Socket client error: ", err.Error())
	}

	// Remove the client from the map of connections
	s.removeClient(&client)

}

// addClient assigns a unique key to a new client, and adds it to the map of connected clients
func (s *SocketsService) addClient(client *SocketConn) error {

	// Control access to the clients map
	s.clientsMut.Lock()
	defer s.clientsMut.Unlock()

	// Start with a random number and increment until we have a unique key
	key := uint64(randKeySource.Int63())
	for {
		_, taken := s.clients[key]
		if !taken {
			break
		}
		key++
	}

	// Assign the key to the client
	client.key = key
	s.clients[key] = client

	// Return without error
	return nil

}

// removeClient is called when a socket connection disconnects from the server
func (s *SocketsService) removeClient(client *SocketConn) {

	// Control access to the clients map
	s.clientsMut.Lock()
	defer s.clientsMut.Unlock()

	// Delete the connection from the map
	delete(s.clients, client.key)

}
