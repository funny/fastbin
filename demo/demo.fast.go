package demo

import (
	"github.com/funny/binary"
)


func (s *Test1) BinarySize() (n int) {
	n = 0 + len(s.Field1) * 8 + len(s.Field2) + len(s.Field3) + s.Field4.BinarySize()	
	for i := 0; i < len(s.Field5); i ++ {
		n += s.Field5[i].BinarySize()
	}
	return
}

func (s *Test1) MarshalBinary() (data []byte, err error) {
	data = make([]byte, s.BinarySize())
	s.MarshalBuffer(&binary.Buffer{Data:data})
	return
}

func (s *Test1) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data:data})
	return nil
}

func (s *Test1) MarshalBuffer(buf *binary.Buffer) {
	buf.WriteUint16LE(uint16(len(s.Field1)))
	for i := 0; i < len(s.Field1); i ++ {
		buf.WriteInt64LE(int64(s.Field1[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field2)))
	buf.WriteBytes(s.Field2)
	buf.WriteUint16LE(uint16(len(s.Field3)))
	buf.WriteString(s.Field3)
	s.Field4.MarshalBuffer(buf)
	buf.WriteUint16LE(uint16(len(s.Field5)))
	for i := 0; i < len(s.Field5); i ++ {
		s.Field5[i].MarshalBuffer(buf)
	}
}

func (s *Test1) UnmarshalBuffer(buf *binary.Buffer) {
	n := 0
	n = int(buf.ReadUint16LE())
	for i := 0; i < n; i ++ {
		s.Field1[i] = int(buf.ReadInt64LE())
	}
	s.Field2 = buf.ReadBytes(int(buf.ReadUint16LE()))
	s.Field3 = buf.ReadString(int(buf.ReadUint16LE()))
	s.Field4.UnmarshalBuffer(buf)
	n = int(buf.ReadUint16LE())
	for i := 0; i < n; i ++ {
		s.Field5[i].UnmarshalBuffer(buf)
	}
}

func (s *Test2) BinarySize() (n int) {
	n = 0 + 8 + 8 + 8 + len(s.Field4) * 8 + len(s.Field5) + len(s.Field6) + 4 + 8 + int(binary.VarintSize(int64(s.Field9)))	
	for i := 0; i < len(s.Field10); i ++ {
		n += int(binary.UvarintSize(uint64(s.Field10[i])))	}
	return
}

func (s *Test2) MarshalBinary() (data []byte, err error) {
	data = make([]byte, s.BinarySize())
	s.MarshalBuffer(&binary.Buffer{Data:data})
	return
}

func (s *Test2) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data:data})
	return nil
}

func (s *Test2) MarshalBuffer(buf *binary.Buffer) {
	buf.WriteInt64LE(int64(s.Field1))
	buf.WriteUint64LE(uint64(s.Field2))
	buf.WriteUint64LE(s.Field3)
	buf.WriteUint16LE(uint16(len(s.Field4)))
	for i := 0; i < len(s.Field4); i ++ {
		buf.WriteInt64LE(int64(s.Field4[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field5)))
	buf.WriteBytes(s.Field5)
	buf.WriteUint16LE(uint16(len(s.Field6)))
	buf.WriteString(s.Field6)
	buf.WriteFloat32LE(s.Field7)
	buf.WriteFloat64LE(s.Field8)
	buf.WriteVarint(int64(s.Field9))
	buf.WriteUint16LE(uint16(len(s.Field10)))
	for i := 0; i < len(s.Field10); i ++ {
		buf.WriteUvarint(uint64(s.Field10[i]))
	}
}

func (s *Test2) UnmarshalBuffer(buf *binary.Buffer) {
	n := 0
	s.Field1 = int(buf.ReadInt64LE())
	s.Field2 = uint(buf.ReadUint64LE())
	s.Field3 = buf.ReadUint64LE()
	n = int(buf.ReadUint16LE())
	for i := 0; i < n; i ++ {
		s.Field4[i] = int(buf.ReadInt64LE())
	}
	s.Field5 = buf.ReadBytes(int(buf.ReadUint16LE()))
	s.Field6 = buf.ReadString(int(buf.ReadUint16LE()))
	s.Field7 = buf.ReadFloat32LE()
	s.Field8 = buf.ReadFloat64LE()
	s.Field9 = varint(buf.ReadVarint())
	n = int(buf.ReadUint16LE())
	for i := 0; i < n; i ++ {
		s.Field10[i] = uvarint(buf.ReadUvarint())
	}
}

