.PHONY: build run test clean deps fmt vet

BINARY_NAME=todo-api

build:
	go build -o $(BINARY_NAME) ./cmd/api

run: build
	./$(BINARY_NAME)

test:
	go test -v ./tests/...

clean:
	rm -f $(BINARY_NAME)
	dep clean

vet:
	go vet ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...

deps:
	go mod download

.DEFAULT_GOAL := build