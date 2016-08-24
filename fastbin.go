package fastbin

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
)

var types []reflect.Type

func Register(v interface{}) {
	RegisterType(reflect.TypeOf(v))
}

func RegisterType(t reflect.Type) {
	for i := 0; i < len(types); i++ {
		if t == types[i] {
			return
		}
	}
	types = append(types, t)
}

func Types() []reflect.Type {
	return types
}

func GenCode() {
	a := newAnalyzer()
	a.Analyze(types)

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		panic("GOPATH environment variable missing")
	}

	path, err := filepath.Abs(gopath)
	if err != nil {
		panic(err)
	}
	path = filepath.Join(path, "src")

	for _, pkg := range a.Packages {
		saveCode(
			filepath.Join(path, pkg.Path),
			filepath.Base(pkg.Path)+".fastbin.go",
			genPackage(pkg),
		)
	}
}

func saveCode(dir, filename string, code []byte) {
	filename = filepath.Join(dir, filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Create file '%s' failed: %s", filename, err)
	}
	if _, err := file.Write(code); err != nil {
		log.Fatalf("Write file '%s' failed: %s", filename, err)
	}
	file.Close()
}
