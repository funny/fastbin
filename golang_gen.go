package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"text/template"
)

func generateGolang(pkg *packageInfo, byteOrder string) (head, code []byte) {
	needN := false
	tpl := template.Must(template.New("code").Funcs(template.FuncMap{
		"TypeName":      goTypeName,
		"MarshalFunc":   goMarshalFunc,
		"UnmarshalFunc": goUnmarshalFunc,
		"TypeInfo": func(t interface{}) *goTplTypeInfo {
			return goTypeInfo(t, byteOrder)
		},
		"SetNeedN": func() string {
			needN = true
			return ""
		},
		"UnsetNeedN": func() string {
			needN = false
			return ""
		},
		"IsNeedN": func() bool {
			return needN
		},
	}).Parse(goTemplate))

	var bf bytes.Buffer
	err := tpl.Execute(&bf, pkg)
	if err != nil {
		log.Fatalf("Generate code failed: %s", err)
	}

	headStr := "package " + pkg.Name + "\n\n import \"github.com/funny/binary\"\n\n"
	if len(pkg.Services) > 0 {
		headStr += `import "github.com/funny/link"`
	}
	head = []byte(headStr)

	code, err = format.Source(bf.Bytes())
	if err != nil {
		fmt.Print(bf.String())
		log.Fatalf("Could't format source: %s", err)
	}
	code = bytes.Replace(code, []byte("\n\n"), []byte("\n"), -1)
	code = bytes.Replace(code, []byte("n = 0\n"), []byte("\n"), -1)
	code = bytes.Replace(code, []byte("+ 0\n"), []byte("\n"), -1)
	return
}

func goTypeName(t *typeInfo) string {
	if t.IsPoint {
		return "*" + goTypeName(t.Type)
	} else if t.IsArray {
		return fmt.Sprintf("[%s]%s", t.Len, goTypeName(t.Type))
	} else if t.Name == "[]byte" {
		return fmt.Sprintf("[%s]byte", t.Len)
	}
	return t.Name
}

type goTplTypeInfo struct {
	Type      *typeInfo
	Name      string
	ByteOrder string
	I         string
	n         int
}

func goTypeInfo(t interface{}, byteOrder string) *goTplTypeInfo {
	switch t1 := t.(type) {
	case *fieldInfo:
		return &goTplTypeInfo{
			t1.Type, "s." + t1.Name, byteOrder, "i0", 0,
		}
	case *goTplTypeInfo:
		t2 := *t1
		if t2.Type.IsArray {
			t2.Name = fmt.Sprintf("%s[i%d]", t2.Name, t2.n)
			t2.n++
			t2.I = fmt.Sprintf("i%d", t2.n)
		} else if t2.Type.IsPoint {
			t2.Name = "(*" + t2.Name + ")"
		}
		t2.Type = t2.Type.Type
		return &t2
	}
	panic("TypeInfo(): Unsuported Type")
}

func goMarshalFunc(t *goTplTypeInfo) string {
	var buf bytes.Buffer
	switch t.Type.Name {
	case "bool":
		fmt.Fprintf(&buf, "if %s { w.WriteUint8(1) } else { w.WriteUint8(0) }", t.Name)
	case "int8", "uint8", "byte":
		fmt.Fprintf(&buf, "w.WriteUint8(uint8(%s))", t.Name)
	case "int16", "uint16":
		fmt.Fprintf(&buf, "w.WriteUint16%s(uint16(%s))", t.ByteOrder, t.Name)
	case "int32", "uint32":
		fmt.Fprintf(&buf, "w.WriteUint32%s(uint32(%s))", t.ByteOrder, t.Name)
	case "int", "uint", "int64", "uint64":
		fmt.Fprintf(&buf, "w.WriteUint64%s(uint64(%s))", t.ByteOrder, t.Name)
	case "float32":
		fmt.Fprintf(&buf, "w.WriteFloat32%s(%s)", t.ByteOrder, t.Name)
	case "float64":
		fmt.Fprintf(&buf, "w.WriteFloat64%s(%s)", t.ByteOrder, t.Name)
	case "string":
		fmt.Fprintf(&buf, "w.WriteUint16%s(uint16(len(%s))); w.WriteString(%s)", t.ByteOrder, t.Name, t.Name)
	case "[]byte":
		if t.Type.Len == "" {
			fmt.Fprintf(&buf, "w.WriteUint16%s(uint16(len(%s))); w.WriteBytes(%s)", t.ByteOrder, t.Name, t.Name)
		} else {
			fmt.Fprintf(&buf, "w.WriteBytes(%s[:])", t.Name)
		}
	default:
		fmt.Fprintf(&buf, "%s.MarshalWriter(w)", t.Name)
	}
	return buf.String()
}

func goUnmarshalFunc(t *goTplTypeInfo) string {
	var buf bytes.Buffer
	switch t.Type.Name {
	case "bool":
		fmt.Fprintf(&buf, "%s = %s(r.ReadUint8() > 0)", t.Name, t.Type.DefName)
	case "int8", "uint8", "byte":
		fmt.Fprintf(&buf, "%s = %s(r.ReadUint8())", t.Name, t.Type.DefName)
	case "int16", "uint16":
		fmt.Fprintf(&buf, "%s = %s(r.ReadUint16%s())", t.Name, t.Type.DefName, t.ByteOrder)
	case "int32", "uint32":
		fmt.Fprintf(&buf, "%s = %s(r.ReadUint32%s())", t.Name, t.Type.DefName, t.ByteOrder)
	case "int", "uint", "int64", "uint64":
		fmt.Fprintf(&buf, "%s = %s(r.ReadUint64%s())", t.Name, t.Type.DefName, t.ByteOrder)
	case "float32":
		fmt.Fprintf(&buf, "%s = %s(r.ReadFloat32%s())", t.Name, t.Type.DefName, t.ByteOrder)
	case "float64":
		fmt.Fprintf(&buf, "%s = %s(r.ReadFloat64%s())", t.Name, t.Type.DefName, t.ByteOrder)
	case "string":
		fmt.Fprintf(&buf, "%s = %s(r.ReadString(int(r.ReadUint16%s())))", t.Name, t.Type.DefName, t.ByteOrder)
	case "[]byte":
		if t.Type.Len == "" {
			fmt.Fprintf(&buf, "%s = %s(r.ReadBytes(int(r.ReadUint16%s())))", t.Name, t.Type.DefName, t.ByteOrder)
		} else {
			fmt.Fprintf(&buf, "io.ReadFull(r, %s[:])", t.Name, t.Type.Len)
		}
	default:
		fmt.Fprintf(&buf, "%s.UnmarshalReader(r)", t.Name)
	}
	return buf.String()
}
