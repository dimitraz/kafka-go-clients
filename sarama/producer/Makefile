DOCKER_ORG=dimitraz
IMAGE_NAME=sarama-producer
BINARY_NAME=sarama-producer
TAG=2.0.0
GOOS=linux

.phony: docker_release
docker_release: docker_build docker_push

.phony: docker_push
docker_push:
	docker push $(DOCKER_ORG)/$(IMAGE_NAME):$(TAG)

.phony: docker_build 
docker_build: build_binary
	docker build -t $(DOCKER_ORG)/$(IMAGE_NAME):$(TAG) .

.PHONY: build_binary
build_binary:
	GOOS=$(GOOS) go build -o $(BINARY_NAME) main.go