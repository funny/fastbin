package demo
import "github.com/funny/binary"
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
	n = 0 + len(s.Field4) + s.Field5.BinarySize()
	for i := 0; i < len(s.Field1); i++ {
		n += len(s.Field1[i])
	}
	for i := 0; i < 10; i++ {
		n += len(s.Field3[i])
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
		buf.WriteUint16LE(uint16(len(s.Field3[i])))
		buf.WriteString(s.Field3[i])
	}
	buf.WriteBytes(s.Field4[:])
	s.Field5.MarshalBuffer(buf)
}
func (s *Test2) UnmarshalBuffer(buf *binary.Buffer) {
	n := 0
	n = int(buf.ReadUint16LE())
	for i := 0; i < n; i++ {
		s.Field1[i] = buf.ReadString(int(buf.ReadUint16LE()))
	}
	for i := 0; i < 10; i++ {
		s.Field3[i] = buf.ReadString(int(buf.ReadUint16LE()))
	}
	copy(s.Field4[:], buf.Take(11))
	s.Field5.UnmarshalBuffer(buf)
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
	n = 0 + 1 + 1 + 2 + 2 + 4 + 4 + 8 + 8 + 8 + 8 + 8*10 + len(s.Field11) + len(s.Field12) + 8*len(s.Field13) + s.Field15.BinarySize()
	for i := 0; i < len(s.Field16); i++ {
		n += s.Field16[i].BinarySize()
	}
	for i := 0; i < 10; i++ {
		n += s.Field17[i].BinarySize()
	}
	return
}
func (s *Test1) MarshalBuffer(buf *binary.Buffer) {
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
	for i := 0; i < n; i++ {
		s.Field13[i] = buf.ReadIntLE()
	}
	for i := 0; i < 10; i++ {
		s.Field14[i] = buf.ReadIntLE()
	}
	s.Field15.UnmarshalBuffer(buf)
	n = int(buf.ReadUint16LE())
	for i := 0; i < n; i++ {
		s.Field16[i].UnmarshalBuffer(buf)
	}
	for i := 0; i < 10; i++ {
		s.Field17[i].UnmarshalBuffer(buf)
	}
}
