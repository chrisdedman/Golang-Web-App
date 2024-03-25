.PHONY: all test vet fmt lint build run deps

all: test vet fmt lint build run deps

deps:
	@echo "Updating dependencies..."
	@go mod tidy

test:
	@echo "Running tests..."
	@go test -v -race -parallel=4 ./...

vet:
	@echo "Running go vet..."
	@go vet ./...

fmt:
	@echo "Formatting code..."
	@go fmt ./...

lint:
	@echo "Running linter..."
	@revive -formatter friendly -config revive.toml ./...

build:
	@echo "Building server binary..."
	@go build -o bin/server ./cmd/server

run: build
	@echo "Running server..."
	@./bin/server
