.PHONY: build
build:
	go build -o bin/egg cmd/egg/main.go

.PHONY: test
test:
	go test -cover -race ./...

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint run
