package gen

var MakefileE = `build:
	go build -o target/${BINARY_NAME}

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
