package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dropdevrahul/gen/gen"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n  %s targetdir pkgname modname\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %-10s: root of the project dir e.g home/user/dir \n\n", "tar_dir")
		fmt.Fprintf(os.Stderr, "  %-10s: package name only one package will be generated  e.g zeus if the flag, this behaviour can be changed with -t flag  \n\n", "pkg_name")
		fmt.Fprintf(os.Stderr, "  %-10s: module name to be put in go.mod e.g github.com/dropdevrahul/gen \n\n", "mod_name")

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  %s: default value %s  \n    %v\n", f.Name, f.Value, f.Usage) // f.Name, f.Value
		})
	}

	pkgType := flag.String("t", "l", "l means generate a single package with pkg_name(library package) \n    i means generate a package with internal and pkg packages ")

	flag.Parse()

	l := len(flag.Args())
	if l != 3 {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println(pkgType)
	gen.GenerateModule(flag.Args()[0], flag.Args()[1], flag.Args()[2])
}
