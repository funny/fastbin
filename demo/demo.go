//go:generate $GOPATH/bin/fastbin
package main

import (
	"fmt"
	"time"

	"github.com/funny/binary"
)

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
	fmt.Printf("Data: %v\n", ab)

	var buf = &binary.Buffer{Data: make([]byte, ab.BinarySize())}

	ab.MarshalWriter(buf)
	ab = AddressBook{}
	ab.UnmarshalReader(buf)

	fmt.Printf("Check: %v\n", ab)
	fmt.Println("Binary size:", len(buf.Data))

	t1 := time.Now()
	for i := 0; i < 1000000; i++ {
		buf.WritePos = 0
		ab.MarshalWriter(buf)
	}
	fmt.Println("Marshal 1M times:", time.Since(t1).String())

	ab = AddressBook{}
	t2 := time.Now()
	for i := 0; i < 1000000; i++ {
		buf.ReadPos = 0
		ab.UnmarshalReader(buf)
	}
	fmt.Println("Umarshal 1M times:", time.Since(t2).String())
}
