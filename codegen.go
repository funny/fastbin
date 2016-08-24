package fastbin

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

var byteOrder = "LE"

func genPackage(pkg *pkgInfo) (code []byte) {
	tpl := template.Must(
		template.New("fastbin").Funcs(template.FuncMap{
			"Package": func() string {
				return filepath.Base(pkg.Path)
			},
			"ByteOrder": func() string {
				return byteOrder
			},
			"TypeName": func(typ reflect.Type) string {
				return goTypeName(pkg.Path, typ)
			},
			"Fields":       goFields,
			"ToType":       goToType,
			"IsFixedSize":  goIsFixedSize,
			"FixedSize":    goFixedSize,
			"NewTplType":   goNewTplType,
			"NewTplType2":  goNewTplType2,
			"NewTplType3":  goNewTplType3,
			"MarshalField": goMarshalField,
			"UnmarshalField": func(typ *tplType) string {
				return goUnmarshalField(pkg.Path, typ)
			},
		}).Parse(appTemplate),
	)

	var bf bytes.Buffer
	err := tpl.Execute(&bf, pkg)
	if err != nil {
		log.Fatalf("Generate code failed: %s", err)
	}

	code, err = format.Source(bf.Bytes())
	if err != nil {
		fmt.Print(bf.String())
		log.Fatalf("Could't format source: %s", err)
	}

	code = bytes.Replace(code, []byte("\n\n"), []byte("\n"), -1)
	code = bytes.Replace(code, []byte("n = 0\n"), []byte("\n"), -1)
	code = bytes.Replace(code, []byte("+ 0\n"), []byte("\n"), -1)
	code, err = format.Source(code)
	if err != nil {
		fmt.Print(bf.String())
		log.Fatalf("Could't format source: %s", err)
	}
	return
}

func goFields(typ reflect.Type) (fields []reflect.StructField) {
	fields = make([]reflect.StructField, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Tag.Get("fb") == "-" {
			continue
		}
		fields = append(fields, field)
	}
	return
}

func goToType(field reflect.StructField) (typ reflect.Type) {
	fb := field.Tag.Get("fb")
	if fb == "" {
		return field.Type
	}

	switch field.Type.Kind() {
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
		reflect.Float32,
		reflect.Float64:
	default:
		if fb != "-" {
			panic(fmt.Sprintf("Unsupported type convertion on '%s'", field.Name))
		}
	}

	switch fb {
	case "int":
		return reflect.TypeOf(int(0))
	case "int8":
		return reflect.TypeOf(int8(0))
	case "int16":
		return reflect.TypeOf(int16(0))
	case "int32":
		return reflect.TypeOf(int32(0))
	case "int64":
		return reflect.TypeOf(int64(0))
	case "uint":
		return reflect.TypeOf(uint(0))
	case "uint8":
		return reflect.TypeOf(uint8(0))
	case "uint16":
		return reflect.TypeOf(uint16(0))
	case "uint32":
		return reflect.TypeOf(uint32(0))
	case "uint64":
		return reflect.TypeOf(uint64(0))
	case "float32":
		return reflect.TypeOf(float32(0))
	case "float64":
		return reflect.TypeOf(float64(0))
	}
	panic(fmt.Sprintf("type '%s' to type '%s' don't support", field.Type, fb))
}

func goIsFixedSize(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Int8, reflect.Uint8, reflect.Bool:
		return true
	case reflect.Int16, reflect.Uint16:
		return true
	case reflect.Int32, reflect.Uint32, reflect.Float32:
		return true
	case reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64, reflect.Float64:
		return true
	case reflect.Array:
		return goIsFixedSize(typ.Elem())
	}
	return false
}

func goFixedSize(typ reflect.Type) int {
	switch typ.Kind() {
	case reflect.Int8, reflect.Uint8, reflect.Bool:
		return 1
	case reflect.Int16, reflect.Uint16:
		return 2
	case reflect.Int32, reflect.Uint32, reflect.Float32:
		return 4
	case reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64, reflect.Float64:
		return 8
	case reflect.Array:
		return typ.Len() * goFixedSize(typ.Elem())
	}
	return 0
}

type tplType struct {
	Name  string
	typ   reflect.Type
	ctyp  reflect.Type
	N     int
	I     string
	Colon string
	K     string
	V     string
	elem  *tplType
}

func goNewTplType(name string, t reflect.Type, n int) *tplType {
	return &tplType{
		Name: name,
		typ:  t,
		N:    n + 1,
		I:    fmt.Sprintf("i%d", n+1),
		K:    fmt.Sprintf("k%d", n+1),
		V:    fmt.Sprintf("v%d", n+1),
	}
}

func goNewTplType2(name string, t reflect.Type, n int) *tplType {
	return &tplType{
		Name:  name,
		typ:   t,
		N:     n + 1,
		I:     fmt.Sprintf("i%d", n+1),
		Colon: ":",
		K:     fmt.Sprintf("k%d", n+1),
		V:     fmt.Sprintf("v%d", n+1),
	}
}

func goNewTplType3(name string, field reflect.StructField, n int) *tplType {
	return &tplType{
		Name: name,
		typ:  field.Type,
		ctyp: goToType(field),
		N:    n + 1,
		I:    fmt.Sprintf("i%d", n+1),
		K:    fmt.Sprintf("k%d", n+1),
		V:    fmt.Sprintf("v%d", n+1),
	}
}

func (t *tplType) Type() reflect.Type {
	if t.ctyp != nil {
		return t.ctyp
	}
	return t.typ
}

func (t *tplType) Elem() *tplType {
	if t.elem == nil {
		if t.IsPtr() {
			if t.Type().Elem().Kind() == reflect.Struct {
				t.elem = goNewTplType(t.Name, t.Type().Elem(), t.N)
			} else {
				t.elem = goNewTplType(t.PtrName(), t.Type().Elem(), t.N)
			}
		} else if t.IsArray() || t.IsSlice() || t.IsMap() {
			t.elem = goNewTplType(t.IndexName(), t.Type().Elem(), t.N)
		}
	}
	return t.elem
}

func (t *tplType) IsFixedSize() bool {
	return goIsFixedSize(t.Type())
}

func (t *tplType) FixedSize() int {
	return goFixedSize(t.Type())
}

func (t *tplType) IndexName() string {
	if t.Type().Kind() == reflect.Map {
		return fmt.Sprintf("(%s[%s])", t.Name, t.K)
	}
	return fmt.Sprintf("(%s[%s])", t.Name, t.I)
}

func (t *tplType) PtrName() string {
	return fmt.Sprintf("(*%s)", t.Name)
}

func (t *tplType) IsBytes() bool {
	typ := t.Type()
	return (typ.Kind() == reflect.Array || typ.Kind() == reflect.Slice) &&
		typ.Elem().Kind() == reflect.Uint8 && typ.Elem().PkgPath() == ""
}

func (t *tplType) IsArray() bool {
	return t.Type().Kind() == reflect.Array
}

func (t *tplType) IsPtr() bool {
	return t.Type().Kind() == reflect.Ptr
}

func (t *tplType) IsSlice() bool {
	return t.Type().Kind() == reflect.Slice
}

func (t *tplType) IsString() bool {
	return t.Type().Kind() == reflect.String
}

func (t *tplType) IsStruct() bool {
	return t.Type().Kind() == reflect.Struct
}

func (t *tplType) IsMap() bool {
	return t.Type().Kind() == reflect.Map
}

func goTypeName(pkgPath string, typ reflect.Type) string {
	if typ.Kind() == reflect.Slice {
		return "[]" + goTypeName(pkgPath, typ.Elem())
	}
	if typ.Kind() == reflect.Ptr {
		return "*" + goTypeName(pkgPath, typ.Elem())
	}
	if typ.Kind() == reflect.Map {
		return "map[" + goTypeName(pkgPath, typ.Key()) + "]" + goTypeName(pkgPath, typ.Elem())
	}
	if pkgPath == typ.PkgPath() {
		return typ.Name()
	}
	return typ.String()
}

func goMarshalField(t *tplType) string {
	switch t.Type().Kind() {
	case reflect.Bool:
		return fmt.Sprintf("if %s { w.WriteUint8(1) } else { w.WriteUint8(0) }", t.Name)
	case reflect.Int8, reflect.Uint8:
		return fmt.Sprintf("w.WriteUint8(uint8(%s))", t.Name)
	case reflect.Int16, reflect.Uint16:
		return fmt.Sprintf("w.WriteUint16%s(uint16(%s))", byteOrder, t.Name)
	case reflect.Int32, reflect.Uint32:
		return fmt.Sprintf("w.WriteUint32%s(uint32(%s))", byteOrder, t.Name)
	case reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64:
		return fmt.Sprintf("w.WriteUint64%s(uint64(%s))", byteOrder, t.Name)
	case reflect.Float32:
		return fmt.Sprintf("w.WriteFloat32%s(%s)", byteOrder, t.Name)
	case reflect.Float64:
		return fmt.Sprintf("w.WriteFloat64%s(%s)", byteOrder, t.Name)
	case reflect.String:
		return fmt.Sprintf("w.WriteUint16%s(uint16(len(%s))); w.WriteString(%s)", byteOrder, t.Name, t.Name)
	case reflect.Struct:
		return fmt.Sprintf("%s.MarshalWriter(w)", t.Name)
	}
	panic(fmt.Sprintf("unsupported type: %s", t.Type()))
}

func goUnmarshalField(pkgPath string, t *tplType) string {
	typeName := goTypeName(pkgPath, t.typ)
	switch t.Type().Kind() {
	case reflect.Bool:
		return fmt.Sprintf("%s %s= %s(r.ReadUint8() > 0)", t.Name, t.Colon, typeName)
	case reflect.Int8, reflect.Uint8:
		return fmt.Sprintf("%s %s= %s(r.ReadUint8())", t.Name, t.Colon, typeName)
	case reflect.Int16, reflect.Uint16:
		return fmt.Sprintf("%s %s= %s(r.ReadUint16%s())", t.Name, t.Colon, typeName, byteOrder)
	case reflect.Int32, reflect.Uint32:
		return fmt.Sprintf("%s %s= %s(r.ReadUint32%s())", t.Name, t.Colon, typeName, byteOrder)
	case reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64:
		return fmt.Sprintf("%s %s= %s(r.ReadUint64%s())", t.Name, t.Colon, typeName, byteOrder)
	case reflect.Float32:
		return fmt.Sprintf("%s %s= %s(r.ReadFloat32%s())", t.Name, t.Colon, typeName, byteOrder)
	case reflect.Float64:
		return fmt.Sprintf("%s %s= %s(r.ReadFloat64%s())", t.Name, t.Colon, typeName, byteOrder)
	case reflect.String:
		return fmt.Sprintf("%s %s= %s(r.ReadString(int(r.ReadUint16%s())))", t.Name, t.Colon, typeName, byteOrder)
	case reflect.Struct:
		if t.Colon != "" {
			return fmt.Sprintf("var %s %s\n%s.UnmarshalReader(r)", t.Name, typeName, t.Name)
		}
		return fmt.Sprintf("%s.UnmarshalReader(r)", t.Name)
	}
	panic(fmt.Sprintf("unsupported type: %s", t.Type()))
}

func line(s string) string {
	return strings.Replace(
		strings.Replace(s, "\n", "", -1), "\t", "", -1,
	)
}
