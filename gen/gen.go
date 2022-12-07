package gen

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func handlePanicError(err error) {
	log.SetOutput(os.Stdout)
	if err != nil {
		fmt.Sprintf(err.Error())
		log.Fatal(err)
	}
}

// GeneratePackage a new folder as a golang module with the given
// package name
// packName is the default package name
// moduleName is the full name of the module e.g.
// github.com/dropdevrahul/gocache
func GenerateModule(target_dir, packName string, moduleName string) {
	p := filepath.Join(target_dir, packName)

	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(p, os.ModePerm)
		handlePanicError(err)
	} else {
		handlePanicError(errors.New("Path already exists"))
	}

	os.Chdir(p)

	if _, err := os.Stat(packName); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(packName, os.ModePerm)
		handlePanicError(err)
	} else {
		handlePanicError(err)
	}

	f, err := os.Create(filepath.Join(packName, packName+".go"))
	handlePanicError(err)

	defer f.Close()
	f.WriteString(fmt.Sprintf("package %s", packName))

	// create main.go file
	fm, err := os.Create("main.go")
	handlePanicError(err)

	defer fm.Close()

	fm.WriteString("package main")
	fm.WriteString("\r\n")
	fm.WriteString("\r\n")
	fm.WriteString("import (")
	fm.WriteString("\r\n")
	fm.WriteString("\r\n")
	fm.WriteString(")")
	fm.WriteString("\r\n")
	fm.WriteString("\r\n")
	fm.WriteString("func main() {")
	fm.WriteString("\r\n")
	fm.WriteString("\r\n")
	fm.WriteString("}")

	cmd := exec.Command("go", "mod", "init", moduleName)
	_, err = cmd.Output()

	handlePanicError(err)
}
