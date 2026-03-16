.PHONY: test test-race fmt lint vet

test:
	go test ./...

test-race:
	go test -race ./...

fmt:
	gofmt -w .
	goimports -w .

lint:
	golangci-lint run ./...

vet:
	go vet ./...
