package main

import "github.com/funny/binary"

func (s *AddressBook) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}

func (s *AddressBook) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}

func (s *AddressBook) BinarySize() (n int) {
	n = 0 + 2

	for i := 0; i < len(s.Person); i++ {

		n += s.Person[i].BinarySize()

	}

	return
}

func (s *AddressBook) MarshalBuffer(buf *binary.Buffer) {

	buf.WriteUint16LE(uint16(len(s.Person)))

	for i := 0; i < len(s.Person); i++ {
		s.Person[i].MarshalBuffer(buf)
	}

}

func (s *AddressBook) UnmarshalBuffer(buf *binary.Buffer) {
	n := 0

	n = int(buf.ReadUint16LE())
	s.Person = make([]Person, n)

	for i := 0; i < n; i++ {
		s.Person[i].UnmarshalBuffer(buf)
	}

}

func (s *Person) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}

func (s *Person) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}

func (s *Person) BinarySize() (n int) {
	n = 0 + 2 + 4 + 2 + 2 + len(s.Name) + len(s.Email)

	for i := 0; i < len(s.Phone); i++ {

		n += s.Phone[i].BinarySize()

	}

	return
}

func (s *Person) MarshalBuffer(buf *binary.Buffer) {

	buf.WriteUint16LE(uint16(len(s.Name)))
	buf.WriteString(s.Name)

	buf.WriteInt32LE(s.Id)

	buf.WriteUint16LE(uint16(len(s.Email)))
	buf.WriteString(s.Email)

	buf.WriteUint16LE(uint16(len(s.Phone)))

	for i := 0; i < len(s.Phone); i++ {
		s.Phone[i].MarshalBuffer(buf)
	}

}

func (s *Person) UnmarshalBuffer(buf *binary.Buffer) {
	n := 0

	s.Name = buf.ReadString(int(buf.ReadUint16LE()))

	s.Id = buf.ReadInt32LE()

	s.Email = buf.ReadString(int(buf.ReadUint16LE()))

	n = int(buf.ReadUint16LE())
	s.Phone = make([]PhoneNum, n)

	for i := 0; i < n; i++ {
		s.Phone[i].UnmarshalBuffer(buf)
	}

}

func (s *PhoneNum) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}

func (s *PhoneNum) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}

func (s *PhoneNum) BinarySize() (n int) {
	n = 0 + 2 + 4 + len(s.Number)

	return
}

func (s *PhoneNum) MarshalBuffer(buf *binary.Buffer) {

	buf.WriteUint16LE(uint16(len(s.Number)))
	buf.WriteString(s.Number)

	buf.WriteInt32LE(s.Type)

}

func (s *PhoneNum) UnmarshalBuffer(buf *binary.Buffer) {

	s.Number = buf.ReadString(int(buf.ReadUint16LE()))

	s.Type = buf.ReadInt32LE()

}

func (s *Test1) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}

func (s *Test1) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}

func (s *Test1) BinarySize() (n int) {
	n = 0 + 1 + 1 + 1 + 2 + 2 + 4 + 4 + 8 + 8 + 8 + 8 + 2 + 2 + 2 + 8*10 + 2 + len(s.Field11) + len(s.Field12) + 8*len(s.Field13) + s.Field15.BinarySize()

	for i := 0; i < len(s.Field16); i++ {

		n += s.Field16[i].BinarySize()

	}

	for i := 0; i < 10; i++ {

		n += s.Field17[i].BinarySize()

	}

	return
}

func (s *Test1) MarshalBuffer(buf *binary.Buffer) {

	if s.Field0 {
		buf.WriteUint8(1)
	} else {
		buf.WriteUint8(0)
	}

	buf.WriteInt8(s.Field1)

	buf.WriteUint8(s.Field2)

	buf.WriteInt16LE(s.Field3)

	buf.WriteUint16LE(s.Field4)

	buf.WriteInt32LE(s.Field5)

	buf.WriteUint32LE(s.Field6)

	buf.WriteInt64LE(s.Field7)

	buf.WriteUint64LE(s.Field8)

	buf.WriteIntLE(s.Field9)

	buf.WriteUintLE(s.Field10)

	buf.WriteUint16LE(uint16(len(s.Field11)))
	buf.WriteString(s.Field11)

	buf.WriteUint16LE(uint16(len(s.Field12)))
	buf.WriteBytes(s.Field12)

	buf.WriteUint16LE(uint16(len(s.Field13)))

	for i := 0; i < len(s.Field13); i++ {
		buf.WriteIntLE(s.Field13[i])
	}

	for i := 0; i < 10; i++ {
		buf.WriteIntLE(s.Field14[i])
	}

	s.Field15.MarshalBuffer(buf)

	buf.WriteUint16LE(uint16(len(s.Field16)))

	for i := 0; i < len(s.Field16); i++ {
		s.Field16[i].MarshalBuffer(buf)
	}

	for i := 0; i < 10; i++ {
		s.Field17[i].MarshalBuffer(buf)
	}

}

func (s *Test1) UnmarshalBuffer(buf *binary.Buffer) {
	n := 0

	s.Field0 = buf.ReadUint8() > 0

	s.Field1 = buf.ReadInt8()

	s.Field2 = buf.ReadUint8()

	s.Field3 = buf.ReadInt16LE()

	s.Field4 = buf.ReadUint16LE()

	s.Field5 = buf.ReadInt32LE()

	s.Field6 = buf.ReadUint32LE()

	s.Field7 = buf.ReadInt64LE()

	s.Field8 = buf.ReadUint64LE()

	s.Field9 = buf.ReadIntLE()

	s.Field10 = buf.ReadUintLE()

	s.Field11 = buf.ReadString(int(buf.ReadUint16LE()))

	s.Field12 = buf.ReadBytes(int(buf.ReadUint16LE()))

	n = int(buf.ReadUint16LE())
	s.Field13 = make([]int, n)

	for i := 0; i < n; i++ {
		s.Field13[i] = buf.ReadIntLE()
	}

	for i := 0; i < 10; i++ {
		s.Field14[i] = buf.ReadIntLE()
	}

	s.Field15.UnmarshalBuffer(buf)

	n = int(buf.ReadUint16LE())
	s.Field16 = make([]Test2, n)

	for i := 0; i < n; i++ {
		s.Field16[i].UnmarshalBuffer(buf)
	}

	for i := 0; i < 10; i++ {
		s.Field17[i].UnmarshalBuffer(buf)
	}

}

func (s *Test2) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}

func (s *Test2) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}

func (s *Test2) BinarySize() (n int) {
	n = 0 + 2 + 2 + 2 + 1 + 1 + 1 + len(s.Field3)

	for i := 0; i < len(s.Field1); i++ {
		n += len(s.Field1[i])
	}

	for i := 0; i < 10; i++ {
		n += len(s.Field2[i])
	}

	if s.Field4 != nil {
		n += s.Field4.BinarySize()
	}

	for i := 0; i < len(s.Field5); i++ {

		if s.Field5[i] != nil {
			n += s.Field5[i].BinarySize()
		}

	}

	for i := 0; i < len(s.Field6); i++ {
		if s.Field6[i] != nil {
			n += 8
		}
	}

	return
}

func (s *Test2) MarshalBuffer(buf *binary.Buffer) {

	buf.WriteUint16LE(uint16(len(s.Field1)))

	for i := 0; i < len(s.Field1); i++ {
		buf.WriteUint16LE(uint16(len(s.Field1[i])))
		buf.WriteString(s.Field1[i])
	}

	for i := 0; i < 10; i++ {
		buf.WriteUint16LE(uint16(len(s.Field2[i])))
		buf.WriteString(s.Field2[i])
	}

	buf.WriteBytes(s.Field3[:])

	if s.Field4 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		s.Field4.MarshalBuffer(buf)
	}

	buf.WriteUint16LE(uint16(len(s.Field5)))

	for i := 0; i < len(s.Field5); i++ {
		if s.Field5[i] == nil {
			buf.WriteUint8(0)
		} else {
			buf.WriteUint8(1)
			s.Field5[i].MarshalBuffer(buf)
		}
	}

	buf.WriteUint16LE(uint16(len(s.Field6)))

	for i := 0; i < len(s.Field6); i++ {
		if s.Field6[i] == nil {
			buf.WriteUint8(0)
		} else {
			buf.WriteUint8(1)
			buf.WriteIntLE(*s.Field6[i])
		}
	}

}

func (s *Test2) UnmarshalBuffer(buf *binary.Buffer) {
	n := 0

	n = int(buf.ReadUint16LE())
	s.Field1 = make([]string, n)

	for i := 0; i < n; i++ {
		s.Field1[i] = buf.ReadString(int(buf.ReadUint16LE()))
	}

	for i := 0; i < 10; i++ {
		s.Field2[i] = buf.ReadString(int(buf.ReadUint16LE()))
	}

	copy(s.Field3[:], buf.Take(11))

	if buf.ReadUint8() == 1 {
		s.Field4.UnmarshalBuffer(buf)
	}

	n = int(buf.ReadUint16LE())
	s.Field5 = make([]*Test3, n)

	for i := 0; i < n; i++ {
		if buf.ReadUint8() == 1 {
			s.Field5[i].UnmarshalBuffer(buf)
		}
	}

	n = int(buf.ReadUint16LE())
	s.Field6 = make([]*int, n)

	for i := 0; i < n; i++ {
		if buf.ReadUint8() == 1 {
			*s.Field6[i] = buf.ReadIntLE()
		}
	}

}

func (s *Test3) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}

func (s *Test3) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}

func (s *Test3) BinarySize() (n int) {
	n = 0 + 8*10

	return
}

func (s *Test3) MarshalBuffer(buf *binary.Buffer) {

	for i := 0; i < 10; i++ {
		buf.WriteIntLE(s.Field1[i])
	}

}

func (s *Test3) UnmarshalBuffer(buf *binary.Buffer) {

	for i := 0; i < 10; i++ {
		s.Field1[i] = buf.ReadIntLE()
	}

}
