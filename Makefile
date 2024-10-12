.PHONY: build
build:
	go build -o bin/egg cmd/egg/main.go

.PHONY: test
test:
	go test -cover -race ./...
