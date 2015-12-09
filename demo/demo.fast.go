package demo
import "github.com/funny/binary"
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
func (s *Test2) BinarySize() (n int) {
	n = 0 + len(s.Field4)
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
}
