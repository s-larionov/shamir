all: test

build:
	go build -o bin/encrypt cmd/encrypt/main.go
	go build -o bin/decrypt cmd/decrypt/main.go

test: build
	go test ./...
