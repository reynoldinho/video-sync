package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool) // Track all clients
	broadcast = make(chan []byte)              // Broadcast channel
	mutex     sync.Mutex                       // Mutex for client map
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func handleSync(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	// Register new client
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	// Cleanup on disconnect
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
	}()

	// Listen for messages (add error handling in production)
	for {
		_, msg, _ := conn.ReadMessage()
		broadcast <- msg
	}
}

func main() {
	go func() {
		for {
			msg := <-broadcast
			mutex.Lock()
			for client := range clients {
				// send message to client
				_ = client.WriteMessage(websocket.TextMessage, msg)
			}
			mutex.Unlock()
		}
	}()

	http.HandleFunc("/sync", handleSync)
	http.ListenAndServe(":8080", nil)
}
