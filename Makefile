.PHONY: test test-unit test-integration fmt vet

test:
	go test -race -tags=integration ./...

test-unit:
	go test -race ./...

test-integration:
	go test -race -run 'Test' -tags=integration -count=1 ./...

fmt:
	gofmt -w .
	goimports -w .

vet:
	go vet ./...
