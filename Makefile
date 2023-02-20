BINARY_NAME=gogen
TAG=$(shell git describe --abbrev --tags)

build:
	go build -o target/${BINARY_NAME}

build-tag:
	echo Building tag: ${TAG}
	git checkout ${TAG}
	go build -o target/${BINARY_NAME}-${TAG}
	git switch -

lint:
	golangci-lint run
