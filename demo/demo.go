//go:generate $GOPATH/bin/fastbin
package demo

type Test1 struct {
	Field0  bool
	Field1  int8
	Field2  uint8
	Field3  int16
	Field4  uint16
	Field5  int32
	Field6  uint32
	Field7  int64
	Field8  uint64
	Field9  int
	Field10 uint
	Field11 string
	Field12 []byte
	Field13 []int
	Field14 [10]int
	Field15 Test2
	Field16 []Test2
	Field17 [10]Test2
}

type Test2 struct {
	Field1 []string
	Field2 [10]string
	Field3 [11]byte
	Field4 *Test3
	Field5 []*Test3
	Field6 []*int
}

type Test3 struct {
	Field1 [10]int
}
