package main

import (
	"log"
	"os"

	kafka "github.com/Shopify/sarama"
)

// Set up the consumer config
func newConfig() *kafka.Config {
	config := kafka.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = kafka.OffsetNewest

	return config
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// Configure the consumer
	config := newConfig()
	brokers := []string{getEnv("SERVERS", "localhost:9092")}

	// Create the consumer
	consumer, err := kafka.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Could not create consumer: %v\n", err)
	}

	defer consumer.Close()

	// Subscribe to topic and consume messages
	topic := getEnv("TOPIC", "test-topic")
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("Error getting partition list: %v\n", err)
	}

	// Consume from every partition in the partition list
	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, partition, kafka.OffsetOldest)
		if err != nil {
			log.Fatalf("Could not create partition consumer: %v\n", err)
		}

		go func(pc kafka.PartitionConsumer) {
			for {
				select {
				case err := <-pc.Errors():
					log.Fatalf("Partition consumer error: %v\n", err)
				}
			}
		}(pc)

		for msg := range pc.Messages() {
			log.Printf("Message on topic: %s, partition: %v, offset: %v: %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		}
	}
}
