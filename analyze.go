package main

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var svcRegexp = regexp.MustCompile(`^\s*fastbin:\s*service(?:\s*=\s*(\d+))?\s*$`)
var msgRegexp = regexp.MustCompile(`^\s*fastbin:\s*message(?:\s*=\s*(\d+))?\s*$`)

type packageInfo struct {
	Name       string
	ConstTypes map[string]string
	Services   map[string]*serviceInfo
	Messages   map[string]*structInfo
}

type serviceInfo struct {
	ID      string
	Name    string
	Recv    string
	Methods []*methodInfo
}

type methodInfo struct {
	ID   string
	Name string
	Type string
}

type structInfo struct {
	ID      string
	Package string
	Name    string
	Fields  []*fieldInfo
}

type fieldInfo struct {
	Name     string
	Type     *typeInfo
	IsFixLen bool
}

type typeInfo struct {
	Name     string
	DefName  string
	Size     int
	Len      string
	Type     *typeInfo
	IsUnknow bool
	IsPoint  bool
	IsArray  bool
}

func analyzeDir(root string) *packageInfo {
	pkgName, fset, files := parseFiles(root)
	pkgAst, _ := ast.NewPackage(fset, files, nil, nil)
	pkgDoc := doc.New(pkgAst, pkgName, doc.AllDecls)

	var pkgInfo = &packageInfo{
		Name:       pkgName,
		ConstTypes: make(map[string]string),
		Services:   make(map[string]*serviceInfo),
		Messages:   make(map[string]*structInfo),
	}
	analyzeConstTypes(pkgInfo, pkgDoc)
	analyzeMessages(pkgInfo, pkgDoc)
	analyzeServices(pkgInfo, pkgDoc)
	return pkgInfo
}

func parseFiles(root string) (string, *token.FileSet, map[string]*ast.File) {
	var pkgName string
	fset := token.NewFileSet()
	files := make(map[string]*ast.File)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		isKissFile, err := filepath.Match("*.kiss.go", info.Name())
		if err != nil {
			log.Fatal("match *.kiss.go file failed: %s", err)
		}
		if isKissFile {
			return nil
		}
		isGoFile, err := filepath.Match("*.go", info.Name())
		if err != nil {
			log.Fatal("match *.go file failed: %s", err)
		}
		if filepath.Dir(path) == root && isGoFile {
			log.Println("<-", path)
			file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				log.Fatalf("Could't parse file '%s': %s", path, err)
			}
			files[path] = file
			pkgName = file.Name.String()
		}
		return nil
	})
	return pkgName, fset, files
}

// find 'type XXX int'
func analyzeConstTypes(pkgInfo *packageInfo, pkgDoc *doc.Package) {
	for _, t := range pkgDoc.Types {
		if len(t.Consts) > 0 {
			if typeSpce, ok := t.Decl.Specs[0].(*ast.TypeSpec); ok {
				declName := typeSpce.Name.String()
				typeName := typeSpce.Type.(*ast.Ident).Name
				pkgInfo.ConstTypes[declName] = typeName
				log.Printf("\tType '%s' -> '%s'", declName, typeName)
			}
		}
	}
}

// find 'kiss:message'
func analyzeMessages(pkgInfo *packageInfo, pkgDoc *doc.Package) {
	for _, t := range pkgDoc.Types {
		if matches := msgRegexp.FindStringSubmatch(t.Doc); len(matches) > 0 {
			typeSpce, ok := t.Decl.Specs[0].(*ast.TypeSpec)
			if !ok {
				log.Fatalf("Found 'kiss:message' tag on non-struct type '%s'", t.Name)
			}
			structType, ok := typeSpce.Type.(*ast.StructType)
			if !ok {
				log.Fatalf("Found 'kiss:message' tag on non-struct type '%s'", t.Name)
			}
			structInfo := analyzeStruct(pkgInfo, matches[1], t.Name, structType)
			pkgInfo.Messages[t.Name] = structInfo
			if structInfo.ID != "" {
				log.Printf("\tMessage '%s', id = %s", t.Name, structInfo.ID)
			} else {
				log.Printf("\tMessage '%s'", t.Name)
			}
			continue
		}
	}
}

// find 'kiss:service'
func analyzeServices(pkgInfo *packageInfo, pkgDoc *doc.Package) {
	for _, t := range pkgDoc.Types {
		if matches := svcRegexp.FindStringSubmatch(t.Doc); len(matches) > 0 {
			service := &serviceInfo{
				ID:   matches[1],
				Name: t.Name,
			}

			pkgInfo.Services[t.Name] = service

			if service.ID != "" {
				log.Printf("\tService '%s', id = %s", service.Name, service.ID)
			} else {
				log.Printf("\tService '%s'", service.Name)
			}
			for _, m := range t.Methods {
				msg := analyzeMethod(pkgInfo, m)
				if msg == nil {
					continue
				}

				// check receiver
				if service.Recv == "" {
					service.Recv = m.Recv
				} else if m.Recv != service.Recv {
					log.Fatalf(
						"Bad service method '%s.%s', service type is '%s', method receiver is '%s'",
						t.Name, m.Name, service.Recv, m.Recv)
				}

				service.Methods = append(service.Methods, &methodInfo{
					ID:   msg.ID,
					Name: m.Name,
					Type: msg.Name,
				})
			}
		}
	}
}

func analyzeMethod(pkgInfo *packageInfo, m *doc.Func) (msg *structInfo) {
	// check parameter number
	params := m.Decl.Type.Params
	if params == nil || len(params.List) != 2 {
		log.Printf("\t\tIgnore method '%s', parameter number != 2", m.Name)
		return
	}

	// first parameter must kiss.Me
	if arg, ok := params.List[0].Type.(*ast.SelectorExpr); !ok || arg.Sel.Name != "Me" {
		log.Printf("\t\tIgnore method '%s', first parameter not kiss.Me", m.Name)
		return
	}

	// match message type
	if arg, ok := params.List[1].Type.(*ast.StarExpr); ok {
		if argType, ok := arg.X.(*ast.Ident); ok {
			msg, ok = pkgInfo.Messages[argType.Name]
			if ok {
				if msg.ID != "" {
					log.Printf("\t\tMethod '%s', id = %s", m.Name, msg.ID)
				} else {
					log.Printf("\t\tMethod '%s'", m.Name)
				}
				return
			}
			log.Printf("\t\tIgnore method '%s', second parameter not a message", m.Name)
			return
		}
		log.Printf("\t\tIgnore method '%s', second parameter not point to an identity", m.Name)
		return
	}
	log.Printf("\t\tIgnore method '%s', second parameter not a pointer", m.Name)
	return
}

func analyzeStruct(pkgInfo *packageInfo, id, structName string, structType *ast.StructType) *structInfo {
	si := &structInfo{ID: id, Name: structName}
	for _, field := range structType.Fields.List {
		typeInfo := analyzeType(pkgInfo, field.Type)
		isFixLen := analyzeFixLen(typeInfo)
		for _, name := range field.Names {
			si.Fields = append(si.Fields, &fieldInfo{
				Name:     name.String(),
				Type:     typeInfo,
				IsFixLen: isFixLen,
			})
		}
	}
	return si
}

func analyzeFixLen(t *typeInfo) bool {
	if t.IsArray {
		if t.Len == "" {
			return false
		}
		return analyzeFixLen(t.Type)
	}
	return t.Size != 0
}

func analyzeType(pkgInfo *packageInfo, astType ast.Expr) *typeInfo {
	var ti typeInfo
	switch t := astType.(type) {
	case *ast.StarExpr:
		ti.IsPoint = true
		ti.Type = analyzeType(pkgInfo, t.X)
	case *ast.ArrayType:
		if size, ok := t.Len.(*ast.BasicLit); ok {
			ti.Len = size.Value
		}
		if t, ok := t.Elt.(*ast.Ident); ok {
			realName := t.Name
			if realName2, ok := pkgInfo.ConstTypes[t.Name]; ok {
				realName = realName2
				if realName == "byte" {
					ti.DefName = t.Name
				}
			}
			if realName == "byte" {
				ti.Name = "[]byte"
				break
			}
		}
		ti.IsArray = true
		ti.Type = analyzeType(pkgInfo, t.Elt)
	case *ast.Ident:
		ti.Name = t.Name
		ti.DefName = ti.Name
		if realName, ok := pkgInfo.ConstTypes[t.Name]; ok {
			ti.DefName = ti.Name
			ti.Name = realName
		}
		switch ti.Name {
		case "string":
			ti.Size = 0
		case "int8", "uint8", "byte", "bool":
			ti.Size = 1
		case "int16", "uint16":
			ti.Size = 2
		case "int32", "uint32", "float32":
			ti.Size = 4
		case "int", "uint", "int64", "uint64", "float64":
			ti.Size = 8
		default:
			ti.IsUnknow = true
		}
	default:
		log.Fatalf("Unsupported field type %#v", astType)
	}
	return &ti
}
