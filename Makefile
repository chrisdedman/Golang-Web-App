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

build: deps
	@echo "Building server binary..."
	@go build -o bin/server ./cmd/server
	@go build -o bin/database ./internal/database

run: build
	@echo "Running server..."
	@./bin/server
	@./bin/database
