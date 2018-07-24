# Kafka Go Client examples

## Sarama
[Sarama](https://github.com/Shopify/sarama) is Shopify's Go client for Kafka and is the most widely adopted. Sarama does not have consumer group support so the [Sarama Cluster](https://github.com/bsm/sarama-cluster) extension was used for this.

```sh
# Start Kafka and Zookeeper
docker-compose up 

# Start the producer, which publishes messages to a given topic every 10 seconds
go run producer/main.go

# Consume the messages
go run consumer/main.go
```