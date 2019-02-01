# Kafka Go Client examples

This repo contains basic examples for setting up simple Kafka producers and consumers in Go.

Examples for the following clients are available:

- [Sarama client](https://github.com/Shopify/sarama) by Shopify in the [sarama](./sarama) folder
- [Go client](https://github.com/confluentinc/confluent-kafka-go) by Confluent in the [confluent](./confluent) folder

## Local development

Use the `docker-compose` file to run Zookeeper and Kafka for local development. Note: this probably won't work for mac.

## Deploying to Kubernetes or Openshift

Use the `deployment.yaml` file to deploy a producer and consumer pod to Openshift. This assumes you're running [Strimzi](http://strimzi.io/) for Kafka and Zookeeper in your Openshift cluster. Instructions for setting up Strimzi can be found in the [Strimzi docs](https://strimzi.io/docs/). Alternatively, if you're interested in playing around with Kafka and Kubeless together, you can get a dev environment started with [this ansible playbook](https://github.com/dimitraz/kafkaless-installer).
