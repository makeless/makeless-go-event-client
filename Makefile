install:
	go get

linter:
	golangci-lint run --timeout 2m0s

test-coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

build:
	go build -o makeless-go-event-client