.PHONY: docker-build

IMAGE_NAME=resu-me_frontend
IMAGE_TAG=$(shell git rev-parse --short HEAD)

docker-build:
	docker build --no-cache -t $(IMAGE_NAME):$(IMAGE_TAG) -t $(IMAGE_NAME):latest .
