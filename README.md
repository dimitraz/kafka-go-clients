# Kafka Go Client examples

This repo contains basic examples for setting up simple Kafka producers and consumers in Go. 

Examples for the following clients are available:
- [Sarama client](https://github.com/Shopify/sarama) by Shopify
- [Go client](https://github.com/confluentinc/confluent-kafka-go) by Confluent 

## Local development 
Use the `docker-compose` file to run Zookeeper and Kafka for local development. Note: this probably won't work for mac.

## Deploying to Kubernetes or Openshift 
Use the `deployment.yaml` file to deploy a producer and consumer pod to Openshift. This assumes you're using [Strimzi](http://strimzi.io/) for Kafka/Zookeeper in the same cluster. Instructions for setting it up can be found in the Strimzi docs. 
