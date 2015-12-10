package main

var goTemplate = `
package {{.Package}}

import "github.com/funny/binary"

{{range .Structs}}
{{NoNeedN}}
func (s *{{.Name}}) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}

func (s *{{.Name}}) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data:data})
	return nil
}

func (s *{{.Name}}) BinarySize() (n int) {
	{{range .Fields}}
		{{template "TypeSize" (TypeInfo .)}}
	{{end}}
	return
}

func (s *{{.Name}}) MarshalBuffer(buf *binary.Buffer) {
	{{range .Fields}}
		{{template "Marshal" (TypeInfo .)}}
	{{end}}
}

func (s *{{.Name}}) UnmarshalBuffer(buf *binary.Buffer) {
	{{DeclN}}
	{{range .Fields}}
		{{template "Unmarshal" (TypeInfo .)}}
	{{end}}
}
{{end}}

{{define "TypeSize"}}
	{{if .Type.IsArray}}
		{{if not .Type.Len}}
		n += 2
		{{NeedN}}
		{{end}}
		for {{.i}} := 0; {{.i}}< {{if .Type.Len}}{{.Type.Len}}{{else}}len({{.Name}}){{end}}; {{.i}}++ {
			{{template "TypeSize" (TypeInfo .)}}
		}
	{{else if .Type.IsPoint}}
		n += 1
		if {{.Name}} != nil {
			{{template "TypeSize" (TypeInfo .)}}
		}
	{{else if .Type.IsUnknow}}
		n += {{.Name}}.BinarySize()
	{{else if or (eq .Type.Name "string") (eq .Type.Name "[]byte") }}
		{{if not .Type.Len}}
		n += 2
		{{end}}
		n += len({{.Name}})
	{{else}}
		n += {{.Type.Size}}
	{{end}}
{{end}}

{{define "Marshal"}}
	{{if .Type.IsArray}}
		{{if not .Type.Len}}buf.WriteUint16LE(uint16(len({{.Name}}))){{end}}
		for {{.i}} := 0; {{.i}}< {{if .Type.Len}}{{.Type.Len}}{{else}}len({{.Name}}){{end}}; {{.i}}++ {
			{{template "Marshal" (TypeInfo .)}}
		}
	{{else if .Type.IsPoint}}
		if {{.Name}} == nil { 
			buf.WriteUint8(0);
		} else {
			buf.WriteUint8(1);
			{{if .Type.Type.IsUnknow}}
				{{.Name}}.MarshalBuffer(buf)
			{{else}}
				{{template "Marshal" (TypeInfo .)}}
			{{end}}
		}
	{{else}}
		{{MarshalFunc .}}
	{{end}}
{{end}}

{{define "Unmarshal"}}
	{{if .Type.IsArray}}
		{{if not .Type.Len}}
		n = int(buf.ReadUint16LE())
		{{.Name}} = make({{TypeName .Type}}, n)
		{{end}}
		for {{.i}} := 0; {{.i}}< {{if .Type.Len}}{{.Type.Len}}{{else}}n{{end}}; {{.i}}++ {
			{{template "Unmarshal" (TypeInfo .)}}
		}
	{{else if .Type.IsPoint}}
		if buf.ReadUint8() == 1 {
			{{if .Type.Type.IsUnknow}}
				{{.Name}} = new({{TypeName .Type.Type}})
				{{.Name}}.UnmarshalBuffer(buf)
			{{else}}
				{{template "Unmarshal" (TypeInfo .)}}
			{{end}}
		}
	{{else}}
		{{UnmarshalFunc .}}
	{{end}}
{{end}}
`
