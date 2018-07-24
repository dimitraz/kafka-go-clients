package main

import (
	"fmt"
	"log"
	"os"
	"time"

	kafka "github.com/Shopify/sarama"
)

// Set up the producer config
func newConfig() *kafka.Config {
	config := kafka.NewConfig()
	config.Producer.RequiredAcks = kafka.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

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
	brokers := []string{getEnv("SERVERS", "localhost:9092")}

	// Create the producer
	producer, err := kafka.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Could not create producer: %v\n", err)
	}

	defer producer.Close()

	// Produce messages to topic
	topic := getEnv("TOPIC", "test-topic")
	for {
		msg := &kafka.ProducerMessage{
			Topic: topic,
			Value: kafka.StringEncoder(fmt.Sprintf("Something Cool on %s", time.Now().Format("Mon Jan 2 15:04:05"))),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Delivery failed to topic: %v, partition: %v, offset: %v. %v\n", topic, partition, offset, err)
		} else {
			log.Printf("Delivered message to topic: %v, partition: %v, offset: %v\n", topic, partition, offset)
		}

		time.Sleep(10 * time.Second)
	}
}
