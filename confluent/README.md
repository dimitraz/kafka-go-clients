# Confluent Go client

Confluent's Go client for Kafka is dependent on the `librdkafka` library which must be installed first to run the examples locally. See instructions [here](https://github.com/confluentinc/confluent-kafka-go#installing-librdkafka).

```sh
# Start Kafka and Zookeeper
docker-compose up

# Start the producer, which publishes messages to a given topic every 10 seconds
go run producer/main.go

# Consume the messages
go run consumer/main.go
```

### Deploying to Kubernetes or Openshift

The `deployment.yaml` file in this repo assumes you're using a [Strimzi](http://strimzi.io/) cluster for Kafka and Zookeeper.

First, build and push the consumer and producer docker images to your dockerhub:

```sh
# Build and push the consumer image to your Dockerhub
export DOCKER_ORG=<your-dockerhub-username>

cd consumer
make docker_release

# Build and push the producer image
cd ../producer
make docker_release
```

---

**Note**: There are a few environment variables in the deployment file that you might want to set, if they differ. These are:

- _SERVERS_: The list of kafka brokers. This defaults to `my-cluster-kafka-bootstrap.strimzi:9092` (`my-cluster` is the name of the cluster, `strimzi` is the namespace where Kafka is running)
- _TOPIC_: The name of the topic to produce to/consume from. Defaults to `test-topic`
- _GROUP_ID_: Consumer group id. Defaults to `kafka-consumer`

---

Update `deployment.yaml` with your dockerhub username, and create the deployments:

```sh
# Deploy to an Openshift cluster
oc create -f deployment.yaml

# Deploy to a k8s cluster
kubectl create -f deployment.yaml
```
