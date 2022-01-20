all: build

DOCKER_IMAGE_NAME=if1nal/otus-hw1
DOCKER_IMAGE_TAG=v0.0.2

.PHONY: build
build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .
	docker push $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)