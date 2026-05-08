.PHONY: docker-build

IMAGE_NAME=resu-me-frontend
IMAGE_TAG=$(shell git rev-parse --short HEAD)

docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -t $(IMAGE_NAME):latest .
