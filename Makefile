.PHONY: all test vet fmt build run deps

all: test vet fmt build run deps

help:
	@echo "Usage: make [command]"
	@echo "  help       - Show this help message"
	@echo "  run        - Start the database, run migrations, and start the application locally"
	@echo "  test       - Run tests on the codebase"
	@echo "  vet        - Run go vet on the codebase to check for errors"
	@echo "  fmt        - Run go fmt on the codebase to format the code"
	@echo "  build      - Build the server binary and database binary"
	@echo "  deps       - Update dependencies using go mod tidy"
	@echo "  clean      - Clean up the project build directory"
	@echo "  all        - Run tests, vet, fmt, build, run, and deps"

deps:
	@echo "Updating dependencies..."
	@go mod tidy

test:
	@echo "Running tests..."
	@go test -v ./internal/...

vet:
	@echo "Running go vet..."
	@go vet ./...

fmt:
	@echo "Formatting code..."
	@go fmt ./...

build: deps
	@echo "Building server binary..."
	@go build -o bin/server ./
	@go build -o bin/database ./config/database

run: vet fmt clean build
	@echo "Running server..."
	@./bin/server
	@./bin/database

clean:
	@echo "Cleaning up..."
	@rm -rf bin