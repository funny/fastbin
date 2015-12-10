package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type File struct {
	Name    string
	Package string
	Structs []*Struct
}

type Struct struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name string
	Type *Type
}

type Type struct {
	Name     string
	Size     int
	Len      string
	Type     *Type
	IsUnknow bool
	IsPoint  bool
	IsArray  bool
}

func analyzeFile(filename string, src interface{}) *File {
	fset := token.NewFileSet()
	pkg, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		log.Fatalf("Could't parse file '%s': %s", filename, err)
	}
	file := File{Name: filename, Package: pkg.Name.String()}
	for _, obj := range pkg.Scope.Objects {
		if obj.Kind == ast.Typ {
			typeSpce, ok := obj.Decl.(*ast.TypeSpec)
			if !ok {
				log.Fatalf("Unknown type without TypeSec: %#v", obj)
			}
			if structType, ok := typeSpce.Type.(*ast.StructType); ok {
				structInfo := analyzeStruct(typeSpce.Name.String(), structType)
				file.Structs = append(file.Structs, structInfo)
			}
		}
	}
	return &file
}

func analyzeStruct(structName string, structType *ast.StructType) *Struct {
	structInfo := Struct{Name: structName}
	for _, field := range structType.Fields.List {
		structInfo.Fields = append(structInfo.Fields, &Field{
			Name: field.Names[0].Name,
			Type: analyzeType(field.Type),
		})
	}
	return &structInfo
}

func analyzeType(astType ast.Expr) *Type {
	var typeInfo Type
	switch t := astType.(type) {
	case *ast.StarExpr:
		typeInfo.Size = 1
		typeInfo.IsPoint = true
		typeInfo.Type = analyzeType(t.X)
	case *ast.ArrayType:
		if size, ok := t.Len.(*ast.BasicLit); ok {
			typeInfo.Len = size.Value
		}
		if t, ok := t.Elt.(*ast.Ident); ok && t.Name == "byte" {
			typeInfo.Size = 1
			typeInfo.Name = "[]byte"
			break
		}
		typeInfo.Size = 2
		typeInfo.IsArray = true
		typeInfo.Type = analyzeType(t.Elt)
	case *ast.Ident:
		typeInfo.Name = t.Name
		switch t.Name {
		case "int8", "uint8", "byte", "bool", "string":
			typeInfo.Size = 1
		case "int16", "uint16":
			typeInfo.Size = 2
		case "int32", "uint32", "float32":
			typeInfo.Size = 4
		case "int", "uint", "int64", "uint64", "float64":
			typeInfo.Size = 8
		default:
			typeInfo.IsUnknow = true
		}
	default:
		log.Fatalf("Unsupported field type %#v", astType)
	}
	return &typeInfo
}
