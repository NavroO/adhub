.PHONY: build test run proto lint

build:
	go build ./...

test:
	go test ./...

run:
	go run ./cmd/ads

proto:
	protoc --go_out=proto .

lint:
	golangci-lint run ./...
