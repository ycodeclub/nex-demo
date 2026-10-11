# Makefile for demo

.PHONY: run build test proto lint

run:
	go run ./cmd/server/...

build:
	go build -o bin/demo ./cmd/server/...

test:
	go test ./... -v -race

proto:
	nex proto

lint:
	golangci-lint run ./...
.DEFAULT_GOAL := run
