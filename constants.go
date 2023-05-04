package gen

var MakefileE = `build:
	go build -o target/${BINARY_NAME %s

lint:
	golangci-lint run

test:
	go test ./...
`

var GitIgnore = "*.swp\nbuilds/\ntarget/"

var MakefileI = `lint:
	golangci-lint run

test:
	go test ./...
`
