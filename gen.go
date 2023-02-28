package gen

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func handlePanicError(err error) {
	log.SetOutput(os.Stdout)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

// GeneratePackage a new folder as a golang module with the given
// package name
// packName is the default package name
// moduleName is the full name of the module e.g.
// github.com/dropdevrahul/gocache
func GenerateModule(target_dir, packName, moduleName, t string) {
	p := filepath.Join(target_dir, packName)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		err := os.Mkdir(p, os.ModePerm)
		handlePanicError(err)
	} else {
		handlePanicError(errors.New("Path already exists"))
	}

	os.Chdir(p)

	f, err := os.Create(packName + ".go")
	handlePanicError(err)

	defer f.Close()
	f.WriteString(fmt.Sprintf("package %s", packName))

	// create main.go file
	if t == "e" {
		err = os.Mkdir("cmd", os.ModePerm)
		handlePanicError(err)

		err = os.Mkdir(filepath.Join("cmd", packName), os.ModePerm)
		handlePanicError(err)

		fm, err := os.Create(filepath.Join("cmd", packName, "main.go"))
		handlePanicError(err)

		defer fm.Close()

		fm.WriteString("package main")
		fm.WriteString("\r\n")
		fm.WriteString("\r\n")
		fm.WriteString("func main() {")
		fm.WriteString("\r\n")
		fm.WriteString("}")
	}

	cmd := exec.Command("go", "mod", "init", moduleName)
	_, err = cmd.Output()
	handlePanicError(err)

	// create commonly used files .gitignore, .golangci.yml
	AddContentsToFile(GitIgnore, ".gitignore")
	AddContentsToFile(GolangCI, ".golangci.yml")

	if t == "e" {
		mk := fmt.Sprintf("BINARY_NAME=%s\n", packName) + MakefileI
		AddContentsToFile(mk, "Makefile")
	} else {
		AddContentsToFile(MakefileE, "Makefile")
	}

	cmd = exec.Command("git", "init")
	err = cmd.Run()
	handlePanicError(err)

	// ssh remote url
	gitUrl := "git@" + strings.Replace(moduleName, ".com/", ".com:", 1) + ".git"
	cmd = exec.Command("git", "remote", "add", "origin", gitUrl)
	err = cmd.Run()
	handlePanicError(err)
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func AddContentsToFile(srcString, dst string) (err error) {
	out, err := os.Create(dst)
	if err != nil {
		handlePanicError(err)

		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	out.Write([]byte(srcString))

	return
}
