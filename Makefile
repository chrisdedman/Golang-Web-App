.PHONY: all test lint build run

all: test lint build run

test:
	go test -v -race -parallel=4 ./...

lint:
	@echo "Running revive..."
	@go list ./... | grep -v /vendor/ | xargs -L1 revive -formatter friendly -config revive.toml

build:
	@echo "Building server binary..."
	@go build -o bin/server ./cmd/server

run: build
	@echo "Running server..."
	@./bin/server
