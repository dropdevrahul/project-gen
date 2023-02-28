package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dropdevrahul/gen"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: %s [OPTIONS] targetdir pkgname modname\n\n", os.Args[0])
		fmt.Fprintf(os.Stdout, "  %-10s: root of the project dir e.g home/user/dir \n\n", "targetdir")
		fmt.Fprintf(os.Stdout, "  %-10s: package name only one package will be generated  e.g zeus if the flag, this behaviour can be changed with -t flag  \n\n", "pkgname")
		fmt.Fprintf(os.Stdout, "  %-10s: module name to be put in go.mod e.g github.com/dropdevrahul/gen \n\n", "modname")

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stdout, "  -%s: default value %s  \n    %v\n", f.Name, f.Value, f.Usage) // f.Name, f.Value
		})
	}

	pkgType := flag.String("t", "l", "l means generate a library package \n    e means generate a executable go package with main.go")

	flag.Parse()

	l := len(flag.Args())
	if l != 3 {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println(pkgType)
	gen.GenerateModule(flag.Args()[0], flag.Args()[1], flag.Args()[2], *pkgType)
}
