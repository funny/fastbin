package main

import "github.com/funny/binary"



	
	


func (s *AddressBook) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *AddressBook) UnmarshalBinary(data []byte) error {
	s.UnmarshalPacket(data)
	return nil
}
func (s *AddressBook) MarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
}
func (s *AddressBook) UnmarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
}
func (s *AddressBook) BinarySize() (n int) {
	
	n += 2
	for i0 := 0; i0 < len(s.Person); i0++ {
		n += s.Person[i0].BinarySize()
	}
	return
}
func (s *AddressBook) MarshalWriter(w binary.BinaryWriter) {
	w.WriteUint16LE(uint16(len(s.Person)))
	for i0 := 0; i0 < len(s.Person); i0++ {
		s.Person[i0].MarshalWriter(w)
	}
}
func (s *AddressBook) UnmarshalReader(r binary.BinaryReader) {
	var n int
	n = int(r.ReadUint16LE())
	s.Person = make([]Person, n)
	for i0 := 0; i0 < n; i0++ {
		s.Person[i0].UnmarshalReader(r)
	}
}
func (s *Person) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *Person) UnmarshalBinary(data []byte) error {
	s.UnmarshalPacket(data)
	return nil
}
func (s *Person) MarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
}
func (s *Person) UnmarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
}
func (s *Person) BinarySize() (n int) {
	n = 4 
	n += 2 + len(s.Name)
	n += 2 + len(s.Email)
	n += 2
	for i0 := 0; i0 < len(s.Phone); i0++ {
		n += s.Phone[i0].BinarySize()
	}
	return
}
func (s *Person) MarshalWriter(w binary.BinaryWriter) {
	w.WriteUint16LE(uint16(len(s.Name)))
	w.WriteString(s.Name)
	w.WriteUint32LE(uint32(s.Id))
	w.WriteUint16LE(uint16(len(s.Email)))
	w.WriteString(s.Email)
	w.WriteUint16LE(uint16(len(s.Phone)))
	for i0 := 0; i0 < len(s.Phone); i0++ {
		s.Phone[i0].MarshalWriter(w)
	}
}
func (s *Person) UnmarshalReader(r binary.BinaryReader) {
	var n int
	s.Name = string(r.ReadString(int(r.ReadUint16LE())))
	s.Id = int32(r.ReadUint32LE())
	s.Email = string(r.ReadString(int(r.ReadUint16LE())))
	n = int(r.ReadUint16LE())
	s.Phone = make([]PhoneNum, n)
	for i0 := 0; i0 < n; i0++ {
		s.Phone[i0].UnmarshalReader(r)
	}
}
func (s *PhoneNum) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *PhoneNum) UnmarshalBinary(data []byte) error {
	s.UnmarshalPacket(data)
	return nil
}
func (s *PhoneNum) MarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
}
func (s *PhoneNum) UnmarshalPacket(p []byte) {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
}
func (s *PhoneNum) BinarySize() (n int) {
	n = 4 
	n += 2 + len(s.Number)
	return
}
func (s *PhoneNum) MarshalWriter(w binary.BinaryWriter) {
	w.WriteUint16LE(uint16(len(s.Number)))
	w.WriteString(s.Number)
	w.WriteUint32LE(uint32(s.Type))
}
func (s *PhoneNum) UnmarshalReader(r binary.BinaryReader) {
	s.Number = string(r.ReadString(int(r.ReadUint16LE())))
	s.Type = int32(r.ReadUint32LE())
}






