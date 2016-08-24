package module

type MyStruct struct {
	Field1 BaseStruct
	Field2 []BaseStruct
	Field3 map[Key]int
}

type Key struct {
	A int
	B int
}

type BaseStruct struct {
	Field1  int
	Field2  uint
	Field3  int16
	Field4  uint16
	Field5  int32
	Field6  uint32
	Field7  int64
	Field8  uint64
	Field9  string
	Field10 []byte

	Field11 []int
	Field12 []string
	Field13 [][]byte
	Field14 [][]int
	Field15 [][]string

	Field16 map[int]int
	Field17 map[int]string
	Field18 map[int][]byte

	Field19 map[string]int
	Field20 map[string]string
	Field21 map[string][]byte

	Field22 map[int]*BaseStruct
	Field23 map[string]*BaseStruct

	Field24 *BaseStruct
	Field25 **BaseStruct
	Field26 []*BaseStruct

	Field27 []map[int]*BaseStruct
	Field28 map[int][]*BaseStruct
}
