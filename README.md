# About

A simple generator to generate a go module with a given package and module name

* Adds a Makefile with [golangci-lint](https://github.com/golangci/golangci-lint) as default linter:
```
make lint

make build 

make test
```

* Initializes a git repo with git and ssh based remote url // TODO to provide an option to use ssh/https

### Usage

```
go build

./gen target/path packagename github/path/name

```


