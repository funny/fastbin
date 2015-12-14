package main

import "strings"

func line(s string) string {
	return strings.Replace(
		strings.Replace(s, "\n", "", -1), "\t", "", -1,
	)
}

var goTemplate = `
package {{.Package}}

import "github.com/funny/binary"

{{range .Structs}}
	{{UnsetNeedN}}
func (s *{{.Name}}) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	data = buf.Data[:buf.WritePos]
	return
}

func (s *{{.Name}}) UnmarshalBinary(data []byte) error {
	s.UnmarshalReader(&binary.Buffer{Data:data})
	return nil
}

func (s *{{.Name}}) BinarySize() (n int) {
	n = ` + line(`
	{{range .Fields}}
		{{if .IsFixLen}}
			{{template "FixLenSize" .Type}}
		{{end}}
	{{end}}0`) + `
	{{range .Fields}}
		{{if not .IsFixLen}}
			{{template "TypeSize" (TypeInfo .)}}
		{{end}}
	{{end}}
	return
}

func (s *{{.Name}}) MarshalWriter(w binary.BinaryWriter) {
	{{range .Fields}}
		{{template "Marshal" (TypeInfo .)}}
	{{end}}
}

func (s *{{.Name}}) UnmarshalReader(r binary.BinaryReader) {
	{{if IsNeedN}}
		var n int
	{{end}}
	{{range .Fields}}
		{{template "Unmarshal" (TypeInfo .)}}
	{{end}}
}
{{end}}

` + line(`
{{define "FixLenSize"}}
	{{if .IsArray}}
		{{.Len}} * {{template "FixLenSize" .Type}}
	{{else}}
		{{.Size}} +
	{{end}}
{{end}}
`) + `
{{define "TypeSize"}}
	{{if .Type.IsArray}}
		{{if not .Type.Len}}
			n += 2{{SetNeedN}}
		{{end}}
		for {{.I}} := 0; {{.I}}< {{if .Type.Len}}{{.Type.Len}}{{else}}len({{.Name}}){{end}}; {{.I}}++ {
			{{template "TypeSize" (TypeInfo .)}}
		}
	{{else if .Type.IsPoint}}
		n += 1
		if {{.Name}} != nil {
			{{template "TypeSize" (TypeInfo .)}}
		}
	{{else if .Type.IsUnknow}}
		n += {{.Name}}.BinarySize()
	{{else if or (eq .Type.Name "string") (eq .Type.Name "[]byte")}}
		n += {{if not .Type.Len}}2 + {{end}}len({{.Name}})
	{{end}}
{{end}}

{{define "Marshal"}}
	{{if .Type.IsArray}}
		{{if not .Type.Len}}
			w.WriteUint16LE(uint16(len({{.Name}})))
		{{end}}
		for {{.I}} := 0; {{.I}}< {{if .Type.Len}}{{.Type.Len}}{{else}}len({{.Name}}){{end}}; {{.I}}++ {
			{{template "Marshal" (TypeInfo .)}}
		}
	{{else if .Type.IsPoint}}
		if {{.Name}} == nil { 
			w.WriteUint8(0);
		} else {
			w.WriteUint8(1);
			{{if .Type.Type.IsUnknow}}
				{{.Name}}.MarshalWriter(w)
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
			n = int(r.ReadUint16LE())
			{{.Name}} = make({{TypeName .Type}}, n)
		{{end}}
		for {{.I}} := 0; {{.I}}< {{if .Type.Len}}{{.Type.Len}}{{else}}n{{end}}; {{.I}}++ {
			{{template "Unmarshal" (TypeInfo .)}}
		}
	{{else if .Type.IsPoint}}
		if r.ReadUint8() == 1 {
			{{if .Type.Type.IsUnknow}}
				{{.Name}} = new({{TypeName .Type.Type}})
				{{.Name}}.UnmarshalReader(r)
			{{else}}
				{{template "Unmarshal" (TypeInfo .)}}
			{{end}}
		}
	{{else}}
		{{UnmarshalFunc .}}
	{{end}}
{{end}}
`
