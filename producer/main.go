package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Set up the producer config
func newConfig() *kafka.ConfigMap {
	config := &kafka.ConfigMap{
		"bootstrap.servers": getEnv("BOOTSTRAP_SERVERS", "localhost"),
	}
	return config
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// Configure the producer
	config := newConfig()

	// Create the producer
	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatal("Could not create producer")
	}

	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					log.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			default:
				log.Printf("Unexpected event: %v\n", ev)
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := getEnv("TOPIC", "myTopic")
	for _, word := range []string{"what", "three", "words"} {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries
	producer.Flush(15 * 1000)
}
