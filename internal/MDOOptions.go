package internal

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

var Version = "0.1.0"
var Name = "Mongo Data Objects"

type MDOOptions struct {
	Rebuild bool
}

func MDOProcess(opt *MDOOptions) {

	fmt.Printf(`o   o o-o    o-o  
|\ /| |  \  o   o 
| O | |   O |   | 
|   | |  /  o   o 
o   o o-o    o-o
`)

	fmt.Printf("Mongo Data Objects: Version: %s\n", Version)

	processDir(".", opt)
	//fmt.Printf("\n")
	//
	//if opt.Watch {
	//	watchFiles(opt)
	//}
}

func processDir(sdir string, opt *MDOOptions) {
	files, err := os.ReadDir(sdir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			processDir(sdir+"/"+file.Name(), opt)
		} else {
			fi, err := file.Info()
			if err != nil {
				log.Fatal(err)
			}
			processFile(sdir, fi, opt)

		}
	}

}

func processFile(sdir string, file fs.FileInfo, opt *MDOOptions) {
	//fileType := "text"

	if !strings.HasSuffix(file.Name(), ".mdo") {
		return
	}

	trimmedName := strings.Split(file.Name(), ".")[0]

	destFName := sdir + "/" + trimmedName + ".java"
	sourceFName := sdir + "/" + file.Name()

	fmt.Printf("\nProcess MDO: %s --> %s", sourceFName, destFName)

	inBytes, err := os.ReadFile(sourceFName)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %s: %s", sourceFName, err))
		return
	}

	sfi, err := os.Stat(sourceFName)
	if err != nil {
		fmt.Printf(" -- Unable to stat: %s : %s\n", sourceFName, err)
		return
	}
	dfi, err := os.Stat(destFName)
	if err == nil {
		if sfi.ModTime().Before(dfi.ModTime()) {
			if !opt.Rebuild {
				fmt.Printf("-- Not modified \n")
				return
			}
		}
	}

	lex := NewLexer(string(inBytes), sourceFName)
	parser := NewParser(lex)
	ast := parser.parse()

	//dirSects := strings.Split(sdir, "/")
	//fileSects := strings.Split(file.Name(), ".")

	// _ = os.Remove(destFName)
	dfile, err := os.Create(destFName)
	if err != nil {
		fmt.Printf(" -- Error: %s\n", err)
		return
	}

	r := NewRender(ast)
	str := r.Render()
	_, err = dfile.WriteString(str)
	err = dfile.Close()
	if err != nil {
		fmt.Printf(" -- Error: %s\n", err)
		return
	}

	//if parser.hasErrors() {
	//	fmt.Printf("\n\n*** Errors found in : %s\n", file.Name())
	//	for _, err := range parser.errors {
	//		fmt.Printf("Error: %d:%d %s\n", err.lineNum, err.linePos, err.msg)
	//	}
	//	fmt.Printf("\n\n")
	//}
	//// fmt.Printf("\n")

}
