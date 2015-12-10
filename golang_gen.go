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

func generateGolang(file *File) {
	var bf bytes.Buffer

	needN := false

	funcMap := template.FuncMap{
		"TypeInfo": func(t interface{}) map[string]interface{} {
			switch t1 := t.(type) {
			case *Field:
				return map[string]interface{}{
					"Type":  t1.Type,
					"Name":  "s." + t1.Name,
					"i":     "i",
					"Index": "",
				}
			case map[string]interface{}:
				if t1["Type"].(*Type).IsArray {
					t1["Name"] = "(" + t1["Name"].(string) + "[" + t1["i"].(string) + "])"
					t1["i"] = t1["i"].(string) + "i"
				} else if t1["Type"].(*Type).IsPoint {
					t1["Name"] = "(*" + t1["Name"].(string) + ")"
				}
				t1["Type"] = t1["Type"].(*Type).Type
				return t1
			}
			panic("TypeInfo(): Unsuported Type")
		},
		"TypeName": goTypeName,
		"NeedN": func() string {
			needN = true
			return ""
		},
		"NoNeedN": func() string {
			needN = false
			return ""
		},
		"DeclN": func() string {
			if needN {
				return "var n int"
			}
			return ""
		},
		"MarshalFunc": func(t map[string]interface{}) string {
			ft := t["Type"].(*Type)
			var f = t["Name"].(string)
			var buf bytes.Buffer
			switch ft.Name {
			case "bool":
				fmt.Fprintf(&buf, "if %s { buf.WriteUint8(1) } else { buf.WriteUint8(0) }", f)
			case "int":
				fmt.Fprintf(&buf, "buf.WriteIntLE(%s)", f)
			case "uint":
				fmt.Fprintf(&buf, "buf.WriteUintLE(%s)", f)
			case "int8":
				fmt.Fprintf(&buf, "buf.WriteInt8(%s)", f)
			case "uint8", "byte":
				fmt.Fprintf(&buf, "buf.WriteUint8(%s)", f)
			case "int16":
				fmt.Fprintf(&buf, "buf.WriteInt16LE(%s)", f)
			case "uint16":
				fmt.Fprintf(&buf, "buf.WriteUint16LE(%s)", f)
			case "int32":
				fmt.Fprintf(&buf, "buf.WriteInt32LE(%s)", f)
			case "uint32":
				fmt.Fprintf(&buf, "buf.WriteUint32LE(%s)", f)
			case "int64":
				fmt.Fprintf(&buf, "buf.WriteInt64LE(%s)", f)
			case "uint64":
				fmt.Fprintf(&buf, "buf.WriteUint64LE(%s)", f)
			case "float32":
				fmt.Fprintf(&buf, "buf.WriteFloat32LE(%s)", f)
			case "float64":
				fmt.Fprintf(&buf, "buf.WriteFloat64LE(%s)", f)
			case "string":
				fmt.Fprintf(&buf, "buf.WriteUint16LE(uint16(len(%s)))\nbuf.WriteString(%s)", f, f)
			case "[]byte":
				if ft.Len == "" {
					fmt.Fprintf(&buf, "buf.WriteUint16LE(uint16(len(%s)))\nbuf.WriteBytes(%s)", f, f)
				} else {
					fmt.Fprintf(&buf, "buf.WriteBytes(%s[:])", f)
				}
			default:
				fmt.Fprintf(&buf, "%s.MarshalBuffer(buf)", f)
			}
			return buf.String()
		},
		"UnmarshalFunc": func(t map[string]interface{}) string {
			ft := t["Type"].(*Type)
			var f = t["Name"].(string)
			var buf bytes.Buffer
			switch ft.Name {
			case "bool":
				fmt.Fprintf(&buf, "%s = buf.ReadUint8() > 0", f)
			case "int":
				fmt.Fprintf(&buf, "%s = buf.ReadIntLE()", f)
			case "uint":
				fmt.Fprintf(&buf, "%s = buf.ReadUintLE()", f)
			case "int8":
				fmt.Fprintf(&buf, "%s = buf.ReadInt8()", f)
			case "uint8", "byte":
				fmt.Fprintf(&buf, "%s = buf.ReadUint8()", f)
			case "int16":
				fmt.Fprintf(&buf, "%s = buf.ReadInt16LE()", f)
			case "uint16":
				fmt.Fprintf(&buf, "%s = buf.ReadUint16LE()", f)
			case "int32":
				fmt.Fprintf(&buf, "%s = buf.ReadInt32LE()", f)
			case "uint32":
				fmt.Fprintf(&buf, "%s = buf.ReadUint32LE()", f)
			case "int64":
				fmt.Fprintf(&buf, "%s = buf.ReadInt64LE()", f)
			case "uint64":
				fmt.Fprintf(&buf, "%s = buf.ReadUint64LE()", f)
			case "float32":
				fmt.Fprintf(&buf, "%s = buf.ReadFloat32LE()", f)
			case "float64":
				fmt.Fprintf(&buf, "%s = buf.ReadFloat64LE()", f)
			case "string":
				fmt.Fprintf(&buf, "%s = buf.ReadString(int(buf.ReadUint16LE()))", f)
			case "[]byte":
				if ft.Len == "" {
					fmt.Fprintf(&buf, "%s = buf.ReadBytes(int(buf.ReadUint16LE()))", f)
				} else {
					fmt.Fprintf(&buf, "copy(%s[:], buf.Take(%s))", f, ft.Len)
				}
			default:
				fmt.Fprintf(&buf, "%s.UnmarshalBuffer(buf)", f)
			}
			return buf.String()
		},
	}

	tpl := template.Must(template.New("code").Funcs(funcMap).Parse(goTemplate))
	if err := tpl.Execute(&bf, file); err != nil {
		log.Fatalf("Generate code failed: %s", err)
	}

	code, err := format.Source(bf.Bytes())
	if err != nil {
		fmt.Print(bf.String())
		log.Fatalf("Could't format source: %s", err)
	}

	code = bytes.Replace(code, []byte("\n\n"), []byte("\n"), -1)

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
