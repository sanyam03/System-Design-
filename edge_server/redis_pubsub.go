package edge_server

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"

	"slack-clone/internal/db"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

// Publish messages to Redis Pub/Sub
func PublishToRedis(msg db.Message) {
	data, _ := json.Marshal(msg)
	err := redisClient.Publish(context.Background(), msg.ChannelIDToString(), data).Err()
	if err != nil {
		log.Printf("Failed to publish to Redis: %v", err)
	}
}

// Subscribe to a Redis channel
func SubscribeToRedis(channelID string) <-chan *redis.Message {
	pubsub := redisClient.Subscribe(context.Background(), channelID)
	return pubsub.Channel()
}




 