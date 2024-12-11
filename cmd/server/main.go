package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"slack-clone/edge_server"
	"slack-clone/internal/api"
	"slack-clone/internal/db"
	"slack-clone/kafka"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Initialize Kafka producer
	edge_server.InitKafkaProducer()

	// Start Kafka consumer (asynchronously)
	go kafka.StartKafkaConsumer()

	// Set up Gin router
	router := gin.Default()

	// Channel endpoints
	router.POST("/channels", api.CreateChannel)

	// User endpoints
	router.POST("/users", api.CreateUser)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK") // Respond with a simple "OK"
    })

	// Message endpoints
	router.POST("/messages", api.CreateMessage) // Create a new message
	router.GET("/channels/:channel_id/messages", api.GetMessagesByChannel) // Get messages by channel

	// WebSocket endpoint for real-time communication
	router.GET("/ws", edge_server.HandleConnections)

	// Start the server
	router.Run(":8081")
}