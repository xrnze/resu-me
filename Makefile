.PHONY: build-service build-frontend build run-service

build-service:
	$(MAKE) -C backend docker-build

build-frontend:
	$(MAKE) -C frontend docker-build

build: build-service build-frontend

run-service:
	docker-compose up -d
