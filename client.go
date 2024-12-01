package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Struct to hold the outgoing WebSocket request
type Request struct {
	ID    int `json:"id"`
	Value int `json:"value"`
}

// Struct to hold the incoming WebSocket response
type Response struct {
	ID     int `json:"id"`
	Square int `json:"square"`
}

func main() {
	requestCount := 10000

	// http get requests
	{
		baseURL := "http://localhost:200/http?value="
		// Start the timer
		startTime := time.Now()

		// Create an HTTP client
		client := &http.Client{}

		for i := 0; i < requestCount; i++ {
			url := fmt.Sprintf("%s%d", baseURL, i)
			// Send GET request
			resp, _ := client.Get(url)
			resp.Body.Close()
		}

		// Measure the time elapsed
		duration := time.Since(startTime)
		fmt.Printf("Process %d HTTP requests in %v\n", requestCount, duration)
	}

	// ws get requests
	{
		conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:200/ws", nil)
		if err != nil {
			log.Fatalf("Failed to connect to WebSocket server: %v\n", err)
		}
		defer conn.Close()

		// Start the timer to measure the time taken
		startTime := time.Now()

		// Send requests one by one and wait for each response
		for i := 0; i < requestCount; i++ {
			// Prepare the request
			req := Request{
				ID:    i,
				Value: i, // Using the counter as the value for simplicity
			}

			// Send the request to the server
			conn.WriteJSON(req)

			// Wait for the response from the server
			var resp Response
			conn.ReadJSON(&resp)
		}

		// Measure the time elapsed
		duration := time.Since(startTime)
		fmt.Printf("Processed %d WS requests in %v\n", requestCount, duration)
	}
}