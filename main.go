package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	var filename string
	if len(flag.Args()) > 0 {
		filename = flag.Arg(0)
	} else {
		filename = os.Getenv("GOFILE")
	}
	file := analyzeFile(filename, nil)
	generateGolang(file)
}
