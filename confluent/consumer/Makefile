DOCKER_ORG=dimitraz
IMAGE_NAME=go-consumer
BINARY_NAME=consumer-bin
TAG=latest
GOOS=linux

.phony: docker_release
docker_release: docker_build docker_push

.phony: docker_push
docker_push:
	docker push $(DOCKER_ORG)/$(IMAGE_NAME):$(TAG)

.phony: docker_build 
docker_build: 
	docker build -t $(DOCKER_ORG)/$(IMAGE_NAME):$(TAG) .

.phony: build
build:
	GOOS=$(GOOS) go build -o $(BINARY_NAME) main.go