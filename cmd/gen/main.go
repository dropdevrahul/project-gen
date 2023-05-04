package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dropdevrahul/gen"
)

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: %s [OPTIONS] modname\n\n", os.Args[0])
		fmt.Fprintf(os.Stdout,
			"%-10s: module name to be put in go.mod e.g github.com/dropdevrahul/gen \n\n", "modname")

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stdout, "  -%s: default value %s  \n    %v\n", f.Name, f.Value, f.Usage) // f.Name, f.Value
		})
	}

	tarDir := flag.String("d", ".",
		fmt.Sprintf("%-10s: root of the project dir e.g home/user/dir", "targetdir"))
	name := flag.String("n", "",
		fmt.Sprintf("%-10s: name of the package however the type of"+
			"pacakge can be changed with the -t flag", "pkgname"))
	pkgType := flag.String("t", "l",
		"l means generate a library package \n e means generate a executable go package with main.go")

	flag.Parse()

	l := len(flag.Args())
	if l != 1 {
		flag.Usage()
		os.Exit(0)
	}

	gen.GenerateModule(*tarDir, *name, flag.Args()[0], *pkgType)
}
