.PHONY: lint test

all:
	make lint
	make test

lint:
	golangci-lint run --config .golangci.yml

test:
	go test -timeout 30s -cover ./...
