package edge_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"slack-clone/internal/db"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handle WebSocket connections
func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer ws.Close()

	for {
		var msg db.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			break
		}

		// Publish message to Redis Pub/Sub
		PublishToRedis(msg)

		// Push message to Kafka
		PublishToKafka(msg)
	}
}
