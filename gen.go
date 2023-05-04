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

// GeneratePackage a new folder as a golang module with the given
// package name
// packName is the default package name
// moduleName is the full name of the module e.g.
// github.com/dropdevrahul/gocache
func GenerateModule(target_dir, packName, moduleName, t string) {
	p := filepath.Join(target_dir, packName)

	log.Printf("Creating path %s", p)

	_, err := os.Stat(p)
	if err != nil {
		// if path does not exist
		if errors.Is(err, os.ErrNotExist) {
			err = os.Mkdir(p, os.ModePerm)
			if err != nil {
				log.Panic(err)
			}
		}
		log.Panic(errors.New("Failed to check existing path"))
	} else {
		log.Panic(errors.New("Path already exists"))
	}

	os.Chdir(p)

	f, err := os.Create(packName + ".go")
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()
	f.WriteString(fmt.Sprintf("package %s", packName))

	// create main.go file
	if t == "e" {
		err = os.Mkdir("cmd", os.ModePerm)
		if err != nil {
			log.Panic(err)
		}

		err = os.Mkdir(filepath.Join("cmd", packName), os.ModePerm)
		if err != nil {
			log.Panic(err)
		}

		fp := filepath.Join("cmd", packName, "main.go")
		fm, err := os.Create(fp)
		if err != nil {
			log.Panic(err)
		}

		fm.WriteString("package main")
		fm.WriteString("\r\n")
		fm.WriteString("\r\n")
		fm.WriteString("func main() {")
		fm.WriteString("\r\n")
		fm.WriteString("}")
		fm.Close()

		mk := fmt.Sprintf("BINARY_NAME=%s\n", packName) + fmt.Sprintf(MakefileE, fp)
		AddContentsToFile(mk, "Makefile")
	} else {
		AddContentsToFile(MakefileI, "Makefile")
	}

	cmd := exec.Command("go", "mod", "init", moduleName)
	_, err = cmd.Output()
	if err != nil {
		log.Print(moduleName)
		log.Panic(err)
	}

	// create commonly used files .gitignore, .golangci.yml
	AddContentsToFile(GitIgnore, ".gitignore")
	AddContentsToFile(GolangCI, ".golangci.yml")

	cmd = exec.Command("git", "init")
	err = cmd.Run()
	if err != nil {
		log.Panic(err)
	}

	// ssh remote url
	gitUrl := "git@" + strings.Replace(moduleName, ".com/", ".com:", 1) + ".git"
	cmd = exec.Command("git", "remote", "add", "origin", gitUrl)
	err = cmd.Run()
	if err != nil {
		log.Panic(err)
	}
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func AddContentsToFile(srcString, dst string) (err error) {
	out, err := os.Create(dst)
	if err != nil {
		log.Println(err)
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
