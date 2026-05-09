IMAGE_NAME=resu-me_backend
IMAGE_TAG=$(shell git rev-parse --short HEAD)

.PHONY: test test-verbose test-coverage test-coverage-html build run vet fmt lint deps clean all docker-build

BINARY_NAME=resume-analyzer
GO_FILES=$(shell find . -name "*.go" -not -path "./vendor/*")

test:
	go test ./... -count=1

test-verbose:
	go test ./... -v -count=1

test-coverage:
	go test ./... -count=1 -coverprofile=coverage.out
	go tool cover -func=coverage.out

test-coverage-html:
	go test ./... -count=1 -coverprofile=coverage.out
	go tool cover -html=coverage.out

build:
	go build -o $(BINARY_NAME) .

run:
	go run .

vet:
	go vet ./...

fmt:
	go fmt ./...

lint: vet
	@test -z $(shell gofmt -l $(GO_FILES)) || (echo "Unformatted files:"; gofmt -l $(GO_FILES); exit 1)

deps:
	go mod tidy
	go mod download

clean:
	rm -f $(BINARY_NAME) coverage.out

all: deps vet test build

docker-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -t $(IMAGE_NAME):latest .
