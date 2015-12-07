//go:generate $GOPATH/bin/fastbin
package demo

type Test1 struct {
	Field1 []int
	Field2 []byte
	Field3 string
	Field4 *Test2
	Field5 []Test2
}

type Test2 struct {
	Field1 int
	Field2 uint
	Field3 uint64
	Field4 []int
	Field5 []byte
	Field6 string
	Field7 float32
	Field8 float64
}
