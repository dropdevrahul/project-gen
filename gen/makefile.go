package gen

var Makefile = `build:
	go build -o target/${BINARY_NAME}

lint:
	golangci-lint run

test:
	go test ./...
`
