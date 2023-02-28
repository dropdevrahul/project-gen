BINARY_NAME=gogen
TAG=$(shell git describe --abbrev --tags)

build:
	go build -o  target/${BINARY_NAME} cmd/gen/main.go

build-tag:
	echo Building tag: ${TAG}
	git checkout ${TAG}
	go build cmd/main.go -o target/${BINARY_NAME}-${TAG}
	git switch -

lint:
	golangci-lint run
