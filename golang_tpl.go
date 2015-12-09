package main

import (
	"strings"
)

var go_template = `
package {{.Package}}

import "github.com/funny/binary"

{{range .Structs}}

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
` + fuck(`
	n = 0
	{{range .Fields}}
		{{if .IsPointer}}
			+ 1
		{{else if and .Size (not .IsArray)}}
			+ {{.Size}}
		{{else if and .Size .ArraySize}}
			+ {{.Size}} * {{.ArraySize}}
		{{else if or (eq .Type "string") (eq .Type "[]byte") (and .IsArray (not .ArraySize))}}
			+ 2
		{{end}}
	{{end}}
	{{range .Fields}}{{if not .IsPointer}}
		{{if and .Size .IsArray (not .ArraySize)}}
			+ {{.Size}} * len(s.{{.Name}})
		{{else if not .IsArray}} 
			{{if or (eq .Type "string") (eq .Type "[]byte")}}
				+ len(s.{{.Name}})
			{{else if .IsUnknow}}
				+ s.{{.Name}}.BinarySize()
			{{end}}
		{{end}}
	{{end}}{{end}}
`) + `
	{{range .Fields}}
		{{if .IsArray}}
			{{if or (eq .Type "string") (eq .Type "[]byte")}}
				for i := 0; i < {{.GoLen}}; i ++ {
					n += len(s.{{.Name}}[i])
				}
			{{else if and .Size .IsPointer}}
				for i := 0; i < {{.GoLen}}; i ++ {
					if s.{{.Name}}[i] != nil {
						n += {{.Size}}
					}
				}
			{{else if .IsUnknow}}
				for i := 0; i < {{.GoLen}}; i ++ {
					{{if .IsPointer}}
						if s.{{.Name}}[i] != nil {
							n += s.{{.Name}}[i].BinarySize()
						}
					{{else}}
						n += s.{{.Name}}[i].BinarySize()
					{{end}}
				}
			{{end}}
		{{else if and .Size .IsPointer}}
			if s.{{.Name}} != nil {
				n += {{.Size}}
			}
		{{else if and .IsUnknow .IsPointer}}
			if s.{{.Name}} != nil {
				n += s.{{.Name}}.BinarySize()
			}
		{{end}}
	{{end}}
	return
}

func (s *{{.Name}}) MarshalBuffer(buf *binary.Buffer) {
	{{range .Fields}}
		{{if .IsArray}}
			{{if not .ArraySize}}
			buf.WriteUint16LE(uint16(len(s.{{.Name}})))
			{{end}}
			for i := 0; i < {{.GoLen}}; i ++ {
				{{.GoEncodeFunc}}
			}
		{{else}}
			{{.GoEncodeFunc}}
		{{end}}
	{{end}}
}

func (s *{{.Name}}) UnmarshalBuffer(buf *binary.Buffer) {
	{{if .GoNeedN}}n := 0{{end}}
	{{range .Fields}}
		{{if .IsArray}}
			{{if not .ArraySize}}
			n = int(buf.ReadUint16LE())
			s.{{.Name}} = make([]{{if .IsPointer}}*{{end}}{{.Type}}, n)
			{{end}}
			for i := 0; i < {{if .ArraySize}}{{.ArraySize}}{{else}}n{{end}}; i ++ {
				{{.GoDecodeFunc}}
			}
		{{else}}
			{{.GoDecodeFunc}}
		{{end}}
	{{end}}
}

{{end}}
`

func fuck(s string) string {
	return strings.Replace(
		strings.Replace(s, "\n", "", -1), "\t", "", -1,
	)
}
