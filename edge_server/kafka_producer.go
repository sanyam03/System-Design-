package edge_server

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"

	"slack-clone/internal/db"
)

var kafkaProducer sarama.SyncProducer

func InitKafkaProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	var err error
	kafkaProducer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
}

// Publish messages to Kafka
func PublishToKafka(msg db.Message) {
	partitionKey := sarama.StringEncoder(msg.ChannelIDToString())
	data, _ := json.Marshal(msg)
	message := &sarama.ProducerMessage{
		Topic: "messages",
		Key:   partitionKey,
		Value: sarama.ByteEncoder(data),
	}

	_, _, err := kafkaProducer.SendMessage(message)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
	}
}



