SHELL := bash
.SHELLFLAGS := -euo pipefail -c

all:
	go build -v ./...

run:
	go run .

test:
	go test ./... -race | grep -v "no test files"

lint:
	golangci-lint run