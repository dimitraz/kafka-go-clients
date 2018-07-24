package main

import (
	"log"
	"os"
	"os/signal"

	kafka "github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// Set up the consumer config
func newConfig() *cluster.Config {
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
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
	groupId := getEnv("GROUP_ID", "kafka-consumer")
	topics := []string{getEnv("TOPIC", "test-topic")}
	brokers := []string{getEnv("SERVERS", "localhost:9092")}

	// Create the consumer
	consumer, err := cluster.NewConsumer(brokers, groupId, topics, config)
	if err != nil {
		log.Fatalf("Could not create consumer: %v\n", err)
	}

	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Fatalf("Error: %v\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				log.Printf("Message on topic: %s, partition: %v, offset: %v: %s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-signals:
			return
		}
	}
}
