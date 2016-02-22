package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"
)

var plugins = make(map[string]func([]*packageInfo))

func main() {
	flag.Parse()

	dir := "."
	if len(flag.Args()) > 0 {
		dir = flag.Arg(0)
	}

	var pkgInfos []*packageInfo
	for _, dir := range strings.Split(dir, ":") {
		absDir, err := filepath.Abs(dir)
		if err != nil {
			log.Fatalf("filepath.Abs(\"%s\") - %s", dir, err)
		}
		log.Print("Analyze ", absDir)
		pkgInfos = append(pkgInfos, analyzeDir(dir))
	}

	plugin := "go"
	if len(flag.Args()) > 0 {
		plugin = flag.Arg(1)
	}
	plugins[plugin](pkgInfos)
}

func line(s string) string {
	return strings.Replace(
		strings.Replace(s, "\n", "", -1), "\t", "", -1,
	)
}
