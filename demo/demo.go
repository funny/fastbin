//go:generate $GOPATH/bin/fastbin
package main

import (
	"github.com/funny/binary"
	"time"
)

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

type AddressBook struct {
	Person []Person
}

type Person struct {
	Name  string
	Id    int32
	Email string
	Phone []PhoneNum
}

type PhoneNum struct {
	Number string
	Type   int32
}

func main() {
	ab := AddressBook{[]Person{
		{"Alice", 10000, "", []PhoneNum{
			{"123456789", 1},
			{"87654321", 2},
		}},
		{"Bob", 20000, "", []PhoneNum{
			{"01234567890", 3},
		}},
	}}

	var buf = &binary.Buffer{Data: make([]byte, ab.BinarySize())}
	println("Size:", len(buf.Data))

	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		buf.WritePos = 0
		ab.MarshalBuffer(buf)
	}
	println("Marshal 1M times:", time.Since(t1).String())

	ab = AddressBook{}
	t2 := time.Now()
	for i := 0; i < 1000000; i++ {
		buf.ReadPos = 0
		ab.UnmarshalBuffer(buf)
	}
	println("Umarshal 1M times:", time.Since(t2).String())
}
