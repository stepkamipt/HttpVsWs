package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"log"

	"github.com/gorilla/websocket"
)

// handler for GET HTTP requests
func httpHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	numberStr := r.URL.Query().Get("number")
	number, _ := strconv.Atoi(numberStr)
	square := number * number
	fmt.Fprintf(w, "stub responce %d\n", square)
}

// handling GET-over-Websocket requests

type Request struct {
	ID    int `json:"id"`
	Value int `json:"value"`
}

type Response struct {
	ID     int `json:"id"`
	Square int `json:"square"`
}

// Create an upgrader to upgrade HTTP to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins, modify this as needed
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Listen for incoming messages
	for {
		// read json
		var req Request
		err := conn.ReadJSON(&req)
		if err != nil {
			log.Printf("Error reading JSON response: %v\n", err)
			break
		}

		// calc responce
		square := req.Value * req.Value

		// create responce
		resp := Response{
			ID:     req.ID,
			Square: square,
		}

		// Send the response back as JSON
		conn.WriteJSON(resp)
	}
}


func main() {
	// Register the handler function for the root URL
	http.HandleFunc("/http", httpHandler)

	http.HandleFunc("/ws", wsHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :200")
	if err := http.ListenAndServe(":200", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}