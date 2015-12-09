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
	Name      string
	Type      string
	Size      string
	Len       string
	IsArray   bool
	ArraySize string
	IsUnknow  bool
	IsPointer bool
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
		// scan ast
		fieldInfo := Field{Name: field.Names[0].Name}
		switch fieldType := field.Type.(type) {
		case *ast.Ident:
			fieldInfo.Type = fieldType.Name
		case *ast.StarExpr:
			fieldInfo.IsPointer = true
			fieldInfo.Type = fieldType.X.(*ast.Ident).Name
		case *ast.ArrayType:
			if size, ok := fieldType.Len.(*ast.BasicLit); ok {
				fieldInfo.ArraySize = size.Value
			}
			switch arrayType := fieldType.Elt.(type) {
			case *ast.Ident:
				if arrayType.Name == "byte" {
					fieldInfo.Type = "[]byte"
				} else {
					fieldInfo.IsArray = true
					fieldInfo.Type = arrayType.Name
				}
			case *ast.StarExpr:
				fieldInfo.IsArray = true
				fieldInfo.IsPointer = true
				fieldInfo.Type = arrayType.X.(*ast.Ident).Name
			default:
				log.Fatalf("Unsupported array type %#v", fieldType.Elt)
			}
		default:
			log.Fatalf("Unsupported field type %#v", field.Type)
		}
		// check field type
		switch fieldInfo.Type {
		case "int8", "uint8", "byte", "bool":
			fieldInfo.Size = "1"
		case "int16", "uint16":
			fieldInfo.Size = "2"
		case "int32", "uint32":
			fieldInfo.Size = "4"
		case "int", "uint", "int64", "uint64":
			fieldInfo.Size = "8"
		case "string", "[]byte":
		default:
			fieldInfo.IsUnknow = true
		}
		structInfo.Fields = append(structInfo.Fields, &fieldInfo)
	}
	return &structInfo
}
