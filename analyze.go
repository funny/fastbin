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

var svcRegexp = regexp.MustCompile(`^\s*fb:service\s*=\s*(\d+)\s*$`)
var hdlRegexp = regexp.MustCompile(`^\s*fb:handler\s*$`)
var msgRegexp = regexp.MustCompile(`^\s*fb:message(?:\s*=\s*(\d+))?\s*$`)

type packageInfo struct {
	fset       *token.FileSet
	Dir        string
	Name       string
	ConstTypes map[string]string
	Files      map[string]*fileInfo
	serviceIDs map[string]string
}

type fileInfo struct {
	Package    string
	FilePath   string
	ServiceID  string
	ConstTypes map[string]string
	Handler    *handlerInfo
	Messages   map[string]*structInfo
}

type handlerInfo struct {
	ID      string
	Name    string
	Recv    string
	Methods map[string]*methodInfo
}

type methodInfo struct {
	ID   string
	Name string
	Type string
}

type structInfo struct {
	ID        string
	ServiceID string
	APIName   string
	Package   string
	Name      string
	Fields    []*fieldInfo
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

func (pkgInfo *packageInfo) File(pos token.Pos) *fileInfo {
	filename := pkgInfo.fset.File(pos).Name()
	file, exists := pkgInfo.Files[filename]
	if !exists {
		file = &fileInfo{
			ServiceID:  pkgInfo.serviceIDs[filename],
			Package:    pkgInfo.Name,
			ConstTypes: pkgInfo.ConstTypes,
			Messages:   make(map[string]*structInfo),
		}
		pkgInfo.Files[filename] = file
	}
	return file
}

func analyzeDir(dir string) *packageInfo {
	pkgName, fset, files, serviceIDs := parseFiles(dir)
	pkgAst, _ := ast.NewPackage(fset, files, nil, nil)
	pkgDoc := doc.New(pkgAst, pkgName, doc.AllDecls)

	var pkgInfo = &packageInfo{
		fset:       fset,
		Dir:        dir,
		Name:       pkgName,
		ConstTypes: make(map[string]string),
		Files:      make(map[string]*fileInfo),
		serviceIDs: serviceIDs,
	}
	analyzeConstTypes(pkgInfo, pkgDoc)
	analyzeMessages(pkgInfo, pkgDoc)
	analyzeHandlers(pkgInfo, pkgDoc)
	return pkgInfo
}

func parseFiles(dir string) (string, *token.FileSet, map[string]*ast.File, map[string]string) {
	var pkgName string
	fset := token.NewFileSet()
	files := make(map[string]*ast.File)
	serviceIDs := make(map[string]string)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		isFbFile, err := filepath.Match("*.fb.go", info.Name())
		if err != nil {
			log.Fatal("match *.fb.go file failed: %s", err)
		}
		if isFbFile {
			return nil
		}
		isGoFile, err := filepath.Match("*.go", info.Name())
		if err != nil {
			log.Fatal("match *.go file failed: %s", err)
		}
		if filepath.Dir(path) == dir && isGoFile {
			log.Println("<-", path)
			file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				log.Fatalf("Could't parse file '%s': %s", path, err)
			}
			_, filename := filepath.Split(path)
			files[filename] = file
			pkgName = file.Name.String()

			// find 'fb:service = 123'
			matches := svcRegexp.FindStringSubmatch(file.Doc.Text())
			if len(matches) > 0 {
				serviceIDs[filename] = matches[1]
			} else {
				serviceIDs[filename] = ""
			}
		}
		return nil
	})
	return pkgName, fset, files, serviceIDs
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

// find 'fb:message'
func analyzeMessages(pkgInfo *packageInfo, pkgDoc *doc.Package) {
	for _, t := range pkgDoc.Types {
		if matches := msgRegexp.FindStringSubmatch(t.Doc); len(matches) > 0 {
			typeSpce, ok := t.Decl.Specs[0].(*ast.TypeSpec)
			if !ok {
				log.Fatalf("Found 'fb:message' tag on non-struct type '%s'", t.Name)
			}
			structType, ok := typeSpce.Type.(*ast.StructType)
			if !ok {
				log.Fatalf("Found 'fb:message' tag on non-struct type '%s'", t.Name)
			}

			structInfo := analyzeStruct(pkgInfo, matches[1], t.Name, structType)
			file := pkgInfo.File(t.Decl.Pos())
			structInfo.ServiceID = file.ServiceID
			file.Messages[t.Name] = structInfo

			if structInfo.ID != "" {
				log.Printf("\t+ Message '%s', ID = %s", t.Name, structInfo.ID)
			} else {
				log.Printf("\t+ Message '%s'", t.Name)
			}
			continue
		}
	}
}

// find 'fb:handler'
func analyzeHandlers(pkgInfo *packageInfo, pkgDoc *doc.Package) {
	for _, t := range pkgDoc.Types {
		if matches := hdlRegexp.FindStringSubmatch(t.Doc); len(matches) > 0 {
			handler := &handlerInfo{
				Name:    t.Name,
				Methods: make(map[string]*methodInfo),
			}

			file := pkgInfo.File(t.Decl.Pos())
			if file.ServiceID == "" {
				log.Fatalf(
					"Found handler '%s' without service tag",
					handler.Name,
				)
			}
			if file.Handler != nil {
				log.Fatalf(
					"Duplicate service handler: %s, %s",
					handler.Name, file.Handler.Name,
				)
			}
			file.Handler = handler
			handler.ID = file.ServiceID
			log.Printf("\t+ Handler '%s'", handler.Name)

			for _, m := range t.Methods {
				msg := analyzeMethod(file, m)
				if msg == nil {
					continue
				}

				// check receiver
				if handler.Recv == "" {
					handler.Recv = m.Recv
				} else if m.Recv != handler.Recv {
					log.Fatalf(
						"Bad handler method '%s.%s', handler type is '%s', method receiver is '%s'",
						t.Name, m.Name, handler.Recv, m.Recv)
				}

				// check duplicate message ID
				if mm, exists := handler.Methods[msg.ID]; exists {
					log.Fatalf(
						"Duplicate message ID '%s' on '%s' for message type '%s', registed message is '%s' and handled by '%s'",
						msg.ID, m.Name, msg.Name, mm.Type, mm.Name,
					)
				}

				msg.APIName = handler.Name + "." + m.Name
				handler.Methods[msg.ID] = &methodInfo{
					ID:   msg.ID,
					Name: m.Name,
					Type: msg.Name,
				}
				log.Printf("\t\t+ Method '%s'", m.Name)
			}
		}
	}
}

func analyzeMethod(file *fileInfo, m *doc.Func) (msg *structInfo) {
	// check parameter number
	params := m.Decl.Type.Params
	if params == nil || len(params.List) != 2 {
		log.Printf("\t\tIgnore method '%s', parameter number != 2", m.Name)
		return
	}

	// first parameter must *link.Session
	arg1Sel, ok := params.List[0].Type.(*ast.SelectorExpr)
	if !ok {
		log.Printf("\t\tIgnore method '%s', first parameter not a selector", m.Name)
		return
	}
	arg1SelX, ok := arg1Sel.X.(*ast.Ident)
	if !ok {
		log.Printf("\t\tIgnore method '%s', first parameter not a package selector", m.Name)
		return
	}
	if arg1SelX.Name != "link" || arg1Sel.Sel.Name != "FbSession" {
		log.Printf("\t\tIgnore method '%s', first parameter *link.Session", m.Name)
		return
	}

	// match message type
	arg2, ok := params.List[1].Type.(*ast.StarExpr)
	if !ok {
		log.Printf("\t\tIgnore method '%s', second parameter not a pointer", m.Name)
		return
	}
	arg2Type, ok := arg2.X.(*ast.Ident)
	if !ok {
		log.Printf("\t\tIgnore method '%s', second parameter not point to an identity", m.Name)
		return
	}
	msg, ok = file.Messages[arg2Type.Name]
	if !ok {
		log.Printf("\t\tIgnore method '%s', second parameter not a message", m.Name)
		return
	}
	if msg.ID == "" {
		log.Printf("\t\tIgnore method '%s', second parameter not have message ID", m.Name)
		return
	}
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
