package fastbin

import (
	"fmt"
	"reflect"
)

type analyzer struct {
	Packages map[string]*pkgInfo
}

type pkgInfo struct {
	Path    string
	Types   map[string]*typeInfo
	Imports []string
}

type typeInfo struct {
	Type reflect.Type
}

func newAnalyzer() *analyzer {
	return &analyzer{
		Packages: map[string]*pkgInfo{},
	}
}

func (this *analyzer) Analyze(types []reflect.Type) {
	for _, typ := range types {
		this.analyzeType(typ)
	}

	for _, pkg := range this.Packages {
		imports := make(map[string]string)
		visited := make(map[reflect.Type]int)

		for _, typ := range pkg.Types {
			this.analyzeImport(imports, visited, pkg.Path, typ.Type)
		}

		pkg.Imports = make([]string, 0, len(imports))
		for _, path := range imports {
			pkg.Imports = append(pkg.Imports, path)
		}
	}
}

func (this *analyzer) getPackage(pkgPath string) *pkgInfo {
	if info, ok := this.Packages[pkgPath]; ok {
		return info
	}
	info := &pkgInfo{
		Path:    pkgPath,
		Types:   map[string]*typeInfo{},
		Imports: []string{},
	}
	this.Packages[pkgPath] = info
	return info
}

func (this *analyzer) analyzeType(typ reflect.Type) {
	switch typ.Kind() {
	case reflect.Struct:
		pkg := this.getPackage(typ.PkgPath())

		if _, ok := pkg.Types[typ.Name()]; ok {
			return
		}

		pkg.Types[typ.Name()] = &typeInfo{typ}

		for i := 0; i < typ.NumField(); i++ {
			this.analyzeType(typ.Field(i).Type)
		}
	case reflect.Ptr:
		this.analyzeType(typ.Elem())
	case reflect.Array:
		this.analyzeType(typ.Elem())
	case reflect.Slice:
		this.analyzeType(typ.Elem())
	case reflect.Map:
		switch typ.Key().Kind() {
		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64,
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
			reflect.String,
			reflect.Struct,
			reflect.Float32,
			reflect.Float64:
			if typ.Key().Kind() == reflect.Struct {
				this.analyzeType(typ.Key())
			}
			if typ.Elem().Kind() == reflect.Struct {
				panic(fmt.Sprintf("unsupported value type '%s' on map '%s'", typ.Key().Kind(), typ))
			}
			this.analyzeType(typ.Elem())
		default:
			panic(fmt.Sprintf("unsupported key type '%s' on map '%s'", typ.Key().Kind(), typ))
		}
	}
}

func (this *analyzer) analyzeImport(imports map[string]string, visited map[reflect.Type]int, pkg string, typ reflect.Type) {
	if _, exists := visited[typ]; exists {
		return
	}
	visited[typ] = 1

L:
	for {
		switch typ.Kind() {
		case reflect.Array,
			reflect.Ptr,
			reflect.Slice,
			reflect.Map:
			typ = typ.Elem()
			if typPkgPath := typ.PkgPath(); typPkgPath != "" && typPkgPath != pkg {
				imports[typPkgPath] = typPkgPath
			}
		default:
			break L
		}
	}

	switch typ.Kind() {
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.String:
		if typPkgPath := typ.PkgPath(); typPkgPath != "" && typPkgPath != pkg {
			imports[typPkgPath] = typPkgPath
		}
	case reflect.Struct:
		if typPkgPath := typ.PkgPath(); typPkgPath == pkg {
			for i := 0; i < typ.NumField(); i++ {
				fieldType := typ.Field(i).Type
				this.analyzeImport(imports, visited, pkg, fieldType)
			}
		}
	}
}
