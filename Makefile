all : test lint build run

test :
	go test ./...

lint :
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build :
	go build -o bin/server ./cmd/server

run :
	go run cmd/server/server.go
