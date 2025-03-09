BINARY_NAME=bdg
GO_FILES=$(wildcard *.go)

build:
	go build -o ${GOBIN}/${BINARY_NAME} ${GO_FILES}

test:
	go test ./...

clean:
	rm -rf ${GOBIN}/${BINARY_NAME}

fmt:
	go fmt ${GO_FILES}

lint:
	go vet ${GO_FILES}

.PHONY: build run test clean fmt lint
