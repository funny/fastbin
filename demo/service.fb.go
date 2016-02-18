package main

import "github.com/funny/binary"

import "github.com/funny/link"

func (this *MyService) ServiceID() byte {
	return 1
}
func (this *MyService) NewRequest(id byte) (link.FbMessage, link.FbHandler) {
	switch id {
	case 1:
		return new(MyMessage1), link.FbHandler(func(s link.FbSession, msg link.FbMessage) {
			this.HandleMessage1(s, msg.(*MyMessage1))
		})
	case 2:
		return new(MyMessage2), link.FbHandler(func(s link.FbSession, msg link.FbMessage) {
			this.HandleMessage2(s, msg.(*MyMessage2))
		})
	}
	panic("*MyService: Unknow Message Type")
}
func (s *MyMessage1) MessageID() byte {
	return 1
}
func (s MyMessage1) ServiceID() byte {
	return 1
}
func (s *MyMessage1) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *MyMessage1) UnmarshalBinary(data []byte) error {
	s.UnmarshalPacket(data)
	return nil
}
func (s *MyMessage1) MarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
}
func (s *MyMessage1) UnmarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
}
func (s *MyMessage1) BinarySize() (n int) {
	
	n += 2 + len(s.Field1)
	n += len(s.Field2) * 8
	return
}
func (s *MyMessage1) MarshalWriter(w binary.BinaryWriter) {
	w.WriteUint16LE(uint16(len(s.Field1)))
	w.WriteBytes(s.Field1)
	w.WriteUint16LE(uint16(len(s.Field2)))
	for i0 := 0; i0 < len(s.Field2); i0++ {
		w.WriteUint64LE(uint64(s.Field2[i0]))
	}
}
func (s *MyMessage1) UnmarshalReader(r binary.BinaryReader) {
	var n int
	s.Field1 = (r.ReadBytes(int(r.ReadUint16LE())))
	n = int(r.ReadUint16LE())
	s.Field2 = make([]int, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field2[i0] = int(r.ReadUint64LE())
	}
}
func (s *MyMessage2) MessageID() byte {
	return 2
}
func (s MyMessage2) ServiceID() byte {
	return 1
}
func (s *MyMessage2) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *MyMessage2) UnmarshalBinary(data []byte) error {
	s.UnmarshalPacket(data)
	return nil
}
func (s *MyMessage2) MarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
}
func (s *MyMessage2) UnmarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
}
func (s *MyMessage2) BinarySize() (n int) {
	n = 8 
	n += 2 + len(s.Field2)
	return
}
func (s *MyMessage2) MarshalWriter(w binary.BinaryWriter) {
	w.WriteUint64LE(uint64(s.Field1))
	w.WriteUint16LE(uint16(len(s.Field2)))
	w.WriteString(s.Field2)
}
func (s *MyMessage2) UnmarshalReader(r binary.BinaryReader) {
	s.Field1 = int(r.ReadUint64LE())
	s.Field2 = string(r.ReadString(int(r.ReadUint16LE())))
}
func (s *MyMessage3) MessageID() byte {
	return 3
}
func (s MyMessage3) ServiceID() byte {
	return 1
}
func (s *MyMessage3) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *MyMessage3) UnmarshalBinary(data []byte) error {
	s.UnmarshalPacket(data)
	return nil
}
func (s *MyMessage3) MarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
}
func (s *MyMessage3) UnmarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
}
func (s *MyMessage3) BinarySize() (n int) {
	n = 8 
	n += 2 + len(s.Field2)
	return
}
func (s *MyMessage3) MarshalWriter(w binary.BinaryWriter) {
	w.WriteUint64LE(uint64(s.Field1))
	w.WriteUint16LE(uint16(len(s.Field2)))
	w.WriteString(s.Field2)
}
func (s *MyMessage3) UnmarshalReader(r binary.BinaryReader) {
	s.Field1 = int(r.ReadUint64LE())
	s.Field2 = string(r.ReadString(int(r.ReadUint16LE())))
}






