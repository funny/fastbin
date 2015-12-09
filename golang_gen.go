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
	var bf bytes.Buffer

	tpl := template.Must(template.New("code").Parse(go_template))
	if err := tpl.Execute(&bf, file); err != nil {
		log.Fatalf("Generate code failed: %s", err)
	}

	code, err := format.Source(bf.Bytes())
	if err != nil {
		fmt.Print(bf.String())
		log.Fatalf("Could't format source: %s", err)
	}

	//code = bytes.Replace(code, []byte("\n\n"), []byte("\n"), -1)

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

func (field *Field) GoLen() string {
	if field.ArraySize == "" {
		return fmt.Sprintf("len(s.%s)", field.Name)
	} else {
		return field.ArraySize
	}
}

func (s *Struct) GoNeedN() bool {
	for _, field := range s.Fields {
		if field.IsArray {
			if field.Size == "" || field.ArraySize == "" {
				return true
			}
		}
	}
	return false
}

func (field *Field) GoEncodeFunc() string {
	f := "s." + field.Name
	if field.IsArray {
		f += "[i]"
	}
	var buf bytes.Buffer
	s := ""
	if field.IsPointer {
		s = "*"
		fmt.Fprintf(&buf, "if %s == nil { buf.WriteUint8(0); } else { buf.WriteUint8(1); ", f)
	}
	switch field.Type {
	case "bool":
		fmt.Fprintf(&buf, "if %s { buf.WriteUint8(1) } else { buf.WriteUint8(0) }", f)
	case "int":
		fmt.Fprintf(&buf, "buf.WriteIntLE(%s%s)", s, f)
	case "uint":
		fmt.Fprintf(&buf, "buf.WriteUintLE(%s%s)", s, f)
	case "int8":
		fmt.Fprintf(&buf, "buf.WriteInt8(%s%s)", s, f)
	case "uint8", "byte":
		fmt.Fprintf(&buf, "buf.WriteUint8(%s%s)", s, f)
	case "int16":
		fmt.Fprintf(&buf, "buf.WriteInt16LE(%s%s)", s, f)
	case "uint16":
		fmt.Fprintf(&buf, "buf.WriteUint16LE(%s%s)", s, f)
	case "int32":
		fmt.Fprintf(&buf, "buf.WriteInt32LE(%s%s)", s, f)
	case "uint32":
		fmt.Fprintf(&buf, "buf.WriteUint32LE(%s%s)", s, f)
	case "int64":
		fmt.Fprintf(&buf, "buf.WriteInt64LE(%s%s)", s, f)
	case "uint64":
		fmt.Fprintf(&buf, "buf.WriteUint64LE(%s%s)", s, f)
	case "string":
		fmt.Fprintf(&buf, "buf.WriteUint16LE(uint16(len(%s)))\nbuf.WriteString(%s)", f, f)
	case "[]byte":
		if field.ArraySize == "" {
			fmt.Fprintf(&buf, "buf.WriteUint16LE(uint16(len(%s)))\nbuf.WriteBytes(%s)", f, f)
		} else {
			fmt.Fprintf(&buf, "buf.WriteBytes(%s[:])", f)
		}
	default:
		fmt.Fprintf(&buf, "%s.MarshalBuffer(buf)", f)
	}
	if field.IsPointer {
		fmt.Fprintf(&buf, " }")
	}
	return buf.String()
}

func (field *Field) GoDecodeFunc() string {
	f := field.Name
	if field.IsArray {
		f += "[i]"
	}
	var buf bytes.Buffer
	s := ""
	if field.IsPointer {
		s = "*"
		fmt.Fprintf(&buf, "if buf.ReadUint8() == 1 { ")
	}
	switch field.Type {
	case "bool":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadUint8() > 0", s, f)
	case "int":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadIntLE()", s, f)
	case "uint":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadUintLE()", s, f)
	case "int8":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadInt8()", s, f)
	case "uint8", "byte":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadUint8()", s, f)
	case "int16":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadInt16LE()", s, f)
	case "uint16":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadUint16LE()", s, f)
	case "int32":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadInt32LE()", s, f)
	case "uint32":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadUint32LE()", s, f)
	case "int64":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadInt64LE()", s, f)
	case "uint64":
		fmt.Fprintf(&buf, "%ss.%s = buf.ReadUint64LE()", s, f)
	case "string":
		fmt.Fprintf(&buf, "s.%s = buf.ReadString(int(buf.ReadUint16LE()))", f)
	case "[]byte":
		if field.ArraySize == "" {
			fmt.Fprintf(&buf, "s.%s = buf.ReadBytes(int(buf.ReadUint16LE()))", f)
		} else {
			fmt.Fprintf(&buf, "copy(s.%s[:], buf.Take(%s))", f, field.ArraySize)
		}
	default:
		fmt.Fprintf(&buf, "s.%s.UnmarshalBuffer(buf)", f)
	}
	if field.IsPointer {
		fmt.Fprintf(&buf, " }")
	}
	return buf.String()
}
