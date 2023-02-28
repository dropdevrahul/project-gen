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
make build

./target/gogen target/path packagename github/path/name

./target/gogen -h

Usage: ./target/gogen [OPTIONS] targetdir pkgname modname

  targetdir : root of the project dir e.g home/user/dir 

  pkgname   : package name only one package will be generated  e.g zeus if the flag, this behaviour can be changed with -t flag  

  modname   : module name to be put in go.mod e.g github.com/dropdevrahul/gen 

  -t: default value l  
    l means generate a library package 
    e means generate a executable go package with main.go

```


