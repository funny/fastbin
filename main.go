package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	byteOrder := "LE"
	if len(flag.Args()) > 0 {
		byteOrder = flag.Arg(0)
	}

	scanDir(".", byteOrder)
}

func scanDir(dir, byteOrder string) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("filepath.Abs(\"%s\") - %s", dir, err)
	}
	log.Print("Analyze ", absDir)
	pkgInfo := analyzeDir(dir)
	for name, file := range pkgInfo.Files {
		if file.Handler != nil || len(file.Messages) > 0 {
			head, code := generateGolang(file, byteOrder)
			save(dir, name[:strings.LastIndex(name, ".")]+".fb.go", head, code)
		}
	}
}

func save(dir, filename string, head, code []byte) {
	log.Println("->", filename)
	filename = filepath.Join(dir, filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Create file '%s' failed: %s", filename, err)
	}
	if _, err := file.Write(head); err != nil {
		log.Fatalf("Write file '%s' failed: %s", filename, err)
	}
	if _, err := file.Write(code); err != nil {
		log.Fatalf("Write file '%s' failed: %s", filename, err)
	}
	file.Close()
}
