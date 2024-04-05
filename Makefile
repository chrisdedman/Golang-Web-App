.PHONY: all test vet fmt build run deps

all: test vet fmt build run deps

help:
	@echo "Available targets:"
	@echo "  make run           	- Start the database, run migrations, and start the application locally"
	@echo "  make test          	- Run tests on the codebase"
	@echo "  make vet           	- Run go vet on the codebase to check for errors"
	@echo "  make fmt           	- Run go fmt on the codebase to format the code"
	@echo "  make build         	- Build the server binary and database binary"
	@echo "  make deps          	- Update dependencies using go mod tidy"
	@echo "  make help          	- Show this help message"

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

build: deps
	@echo "Building server binary..."
	@go build -o bin/server ./config/server
	@go build -o bin/database ./config/database

run: vet fmt build
	@echo "Running server..."
	@./bin/server
	@./bin/database

clean:
	@echo "Cleaning up..."
	@rm -rf bin