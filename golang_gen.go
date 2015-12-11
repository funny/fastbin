package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"
)

func generateGolang(file *File) {
	needN := false

	tpl := template.Must(template.New("code").Funcs(template.FuncMap{
		"TypeName":      goTypeName,
		"TypeInfo":      goTypeInfo,
		"MarshalFunc":   goMarshalFunc,
		"UnmarshalFunc": goUnmarshalFunc,
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

	if err := tpl.Execute(&bf, file); err != nil {
		log.Fatalf("Generate code failed: %s", err)
	}

	code, err := format.Source(bytes.Replace(bf.Bytes(), []byte("\n\n"), []byte("\n"), -1))
	if err != nil {
		fmt.Print(bf.String())
		log.Fatalf("Could't format source: %s", err)
	}

	if len(flag.Args()) == 0 {
		filename := strings.Replace(file.Name, ".go", ".fast.go", 1)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Could't create file '%s': %s", filename, err)
		}
		if _, err := file.Write(code); err != nil {
			log.Fatalf("Write file '%s' failed: %s", filename, err)
		}
		file.Close()
	} else {
		fmt.Print(string(code))
	}
}

func goTypeName(t *Type) string {
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
	Type *Type
	Name string
	I    string
	n    int
}

func goTypeInfo(t interface{}) *goTplTypeInfo {
	switch t1 := t.(type) {
	case *Field:
		return &goTplTypeInfo{
			t1.Type, "s." + t1.Name, "i0", 0,
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
		fmt.Fprintf(&buf, "if %s { buf.WriteUint8(1) } else { buf.WriteUint8(0) }", t.Name)
	case "int":
		fmt.Fprintf(&buf, "buf.WriteIntLE(%s)", t.Name)
	case "uint":
		fmt.Fprintf(&buf, "buf.WriteUintLE(%s)", t.Name)
	case "int8":
		fmt.Fprintf(&buf, "buf.WriteInt8(%s)", t.Name)
	case "uint8", "byte":
		fmt.Fprintf(&buf, "buf.WriteUint8(%s)", t.Name)
	case "int16":
		fmt.Fprintf(&buf, "buf.WriteInt16LE(%s)", t.Name)
	case "uint16":
		fmt.Fprintf(&buf, "buf.WriteUint16LE(%s)", t.Name)
	case "int32":
		fmt.Fprintf(&buf, "buf.WriteInt32LE(%s)", t.Name)
	case "uint32":
		fmt.Fprintf(&buf, "buf.WriteUint32LE(%s)", t.Name)
	case "int64":
		fmt.Fprintf(&buf, "buf.WriteInt64LE(%s)", t.Name)
	case "uint64":
		fmt.Fprintf(&buf, "buf.WriteUint64LE(%s)", t.Name)
	case "float32":
		fmt.Fprintf(&buf, "buf.WriteFloat32LE(%s)", t.Name)
	case "float64":
		fmt.Fprintf(&buf, "buf.WriteFloat64LE(%s)", t.Name)
	case "string":
		fmt.Fprintf(&buf, "buf.WriteUint16LE(uint16(len(%s)))\nbuf.WriteString(%s)", t.Name, t.Name)
	case "[]byte":
		if t.Type.Len == "" {
			fmt.Fprintf(&buf, "buf.WriteUint16LE(uint16(len(%s)))\nbuf.WriteBytes(%s)", t.Name, t.Name)
		} else {
			fmt.Fprintf(&buf, "buf.WriteBytes(%s[:])", t.Name)
		}
	default:
		fmt.Fprintf(&buf, "%s.MarshalBuffer(buf)", t.Name)
	}
	return buf.String()
}

func goUnmarshalFunc(t *goTplTypeInfo) string {
	var buf bytes.Buffer
	switch t.Type.Name {
	case "bool":
		fmt.Fprintf(&buf, "%s = buf.ReadUint8() > 0", t.Name)
	case "int":
		fmt.Fprintf(&buf, "%s = buf.ReadIntLE()", t.Name)
	case "uint":
		fmt.Fprintf(&buf, "%s = buf.ReadUintLE()", t.Name)
	case "int8":
		fmt.Fprintf(&buf, "%s = buf.ReadInt8()", t.Name)
	case "uint8", "byte":
		fmt.Fprintf(&buf, "%s = buf.ReadUint8()", t.Name)
	case "int16":
		fmt.Fprintf(&buf, "%s = buf.ReadInt16LE()", t.Name)
	case "uint16":
		fmt.Fprintf(&buf, "%s = buf.ReadUint16LE()", t.Name)
	case "int32":
		fmt.Fprintf(&buf, "%s = buf.ReadInt32LE()", t.Name)
	case "uint32":
		fmt.Fprintf(&buf, "%s = buf.ReadUint32LE()", t.Name)
	case "int64":
		fmt.Fprintf(&buf, "%s = buf.ReadInt64LE()", t.Name)
	case "uint64":
		fmt.Fprintf(&buf, "%s = buf.ReadUint64LE()", t.Name)
	case "float32":
		fmt.Fprintf(&buf, "%s = buf.ReadFloat32LE()", t.Name)
	case "float64":
		fmt.Fprintf(&buf, "%s = buf.ReadFloat64LE()", t.Name)
	case "string":
		fmt.Fprintf(&buf, "%s = buf.ReadString(int(buf.ReadUint16LE()))", t.Name)
	case "[]byte":
		if t.Type.Len == "" {
			fmt.Fprintf(&buf, "%s = buf.ReadBytes(int(buf.ReadUint16LE()))", t.Name)
		} else {
			fmt.Fprintf(&buf, "copy(%s[:], buf.Take(%s))", t.Name, t.Type.Len)
		}
	default:
		fmt.Fprintf(&buf, "%s.UnmarshalBuffer(buf)", t.Name)
	}
	return buf.String()
}
