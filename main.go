package main

import (
	"os"

	"github.com/dropdevrahul/go-module-generator/generator"
)

func main() {
	generator.GenerateModule(os.Args[1], os.Args[2])
}
