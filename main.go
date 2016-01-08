package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()

	root := "."
	if len(flag.Args()) > 0 {
		root = flag.Arg(0)
	}

	byteOrder := "LE"
	if len(flag.Args()) > 1 {
		byteOrder = flag.Arg(1)
	}

	if root == "..." || root == "./..." {
		switch root {
		case "...":
			root = os.Getenv("GOPATH")
		case "./...":
			root = "."
		}
		filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if info != nil && info.IsDir() {
				log.Println("scan", path)
				scanDir(path, byteOrder)
			}
			return nil
		})
	} else {
		root = filepath.Clean(root)
		scanDir(root, byteOrder)
	}
}

func scanDir(dir, byteOrder string) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("filepath.Abs(\"%s\") - %s", dir, err)
	}
	log.Print("Analyze ", absDir)
	pkgInfo := analyzeDir(dir)
	if len(pkgInfo.Services) > 0 || len(pkgInfo.Messages) > 0 {
		head, code := generateGolang(pkgInfo, byteOrder)
		save(dir, pkgInfo.Name+".fb.go", head, code)
	} else {
		log.Println("Nothing to do")
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
