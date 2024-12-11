package kafka

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"

	"slack-clone/internal/db"
)

func StartKafkaConsumer() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("messages", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		var message db.Message
		json.Unmarshal(msg.Value, &message)

		// Save to database
		if err := db.DB.Create(&message).Error; err != nil {
			log.Printf("Failed to save message to DB: %v", err)
		}
	}
}
