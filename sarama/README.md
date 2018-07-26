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
### Deploying to Kubernetes or Openshift
The `deployment.yaml` file in this repo assumes you're using a [Strimzi](http://strimzi.io/) cluster for Kafka and Zookeeper. 

**Note**: If your Kafka cluster's name is not `my-cluster-kafka`, you need to update the `SERVER` env vars in the deployment with the actual cluster name.

```sh
# Build and push the consumer image to your Dockerhub
export DOCKER_ORG=<your-dockerhub-username>

cd consumer
make docker_release

# Build and push the producer image
cd ../producer
make docker_release
```

Update `deployment.yaml` with your dockerhub username, and create the `Deployment`s:

```sh
# Deploy to an Openshift cluster
oc create -f deployment.yaml

# Deploy to a k8s cluster
kubectl create -f deployment.yaml
```