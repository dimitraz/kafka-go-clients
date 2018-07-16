DOCKER_ORG=dimitraz
IMAGE_NAME=go-producer
BINARY_NAME=producer-bin
GOOS=linux

.phony: docker_release
docker_release: docker_build docker_push

.phony: docker_push
docker_push:
	docker push $(DOCKER_ORG)/$(IMAGE_NAME)

.phony: docker_build 
docker_build: 
	docker build -t $(DOCKER_ORG)/$(IMAGE_NAME) .

.PHONY: build
build:
	GOOS=$(GOOS) go build -o $(BINARY_NAME) main.go