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
	n = 2 + 0
	for i0 := 0; i0 < len(s.Person); i0++ {
		n += (s.Person[i0]).BinarySize()
	}
	return
}
func (s *AddressBook) MarshalBuffer(buf *binary.Buffer) {
	buf.WriteUint16LE(uint16(len(s.Person)))
	for i0 := 0; i0 < len(s.Person); i0++ {
		(s.Person[i0]).MarshalBuffer(buf)
	}
}
func (s *AddressBook) UnmarshalBuffer(buf *binary.Buffer) {
	var n int
	n = int(buf.ReadUint16LE())
	s.Person = make([]Person, n)
	for i0 := 0; i0 < n; i0++ {
		(s.Person[i0]).UnmarshalBuffer(buf)
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
	n = 2 + 4 + 2 + 2 + 0
	n += len(s.Name)
	n += len(s.Email)
	for i0 := 0; i0 < len(s.Phone); i0++ {
		n += (s.Phone[i0]).BinarySize()
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
	for i0 := 0; i0 < len(s.Phone); i0++ {
		(s.Phone[i0]).MarshalBuffer(buf)
	}
}
func (s *Person) UnmarshalBuffer(buf *binary.Buffer) {
	var n int
	s.Name = buf.ReadString(int(buf.ReadUint16LE()))
	s.Id = buf.ReadInt32LE()
	s.Email = buf.ReadString(int(buf.ReadUint16LE()))
	n = int(buf.ReadUint16LE())
	s.Phone = make([]PhoneNum, n)
	for i0 := 0; i0 < n; i0++ {
		(s.Phone[i0]).UnmarshalBuffer(buf)
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
	n = 2 + 4 + 0
	n += len(s.Number)
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
