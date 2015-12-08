//go:generate $GOPATH/bin/fastbin
package demo

type varint int64
type uvarint uint64

type Test1 struct {
	Field1 []int
	Field2 []byte
	Field3 string
	Field4 *Test2
	Field5 []Test2
	Field6 [10]int
	Field7 [10]byte
}

type Test2 struct {
	Field1  int
	Field2  uint
	Field3  uint64
	Field4  []int
	Field5  []byte
	Field6  string
	Field7  float32
	Field8  float64
	Field9  varint
	Field10 []uvarint
}
