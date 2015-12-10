package main
import "github.com/funny/binary"
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
	n += 2
	n += len(s.Number)
	n += 4
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
	n += 2
	for i := 0; i < len(s.Person); i++ {
		n += (s.Person[i]).BinarySize()
	}
	return
}
func (s *AddressBook) MarshalBuffer(buf *binary.Buffer) {
	buf.WriteUint16LE(uint16(len(s.Person)))
	for i := 0; i < len(s.Person); i++ {
		(s.Person[i]).MarshalBuffer(buf)
	}
}
func (s *AddressBook) UnmarshalBuffer(buf *binary.Buffer) {
	var n int
	n = int(buf.ReadUint16LE())
	s.Person = make([]Person, n)
	for i := 0; i < n; i++ {
		(s.Person[i]).UnmarshalBuffer(buf)
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
	n += 2
	n += len(s.Name)
	n += 4
	n += 2
	n += len(s.Email)
	n += 2
	for i := 0; i < len(s.Phone); i++ {
		n += (s.Phone[i]).BinarySize()
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
		(s.Phone[i]).MarshalBuffer(buf)
	}
}
func (s *Person) UnmarshalBuffer(buf *binary.Buffer) {
	var n int
	s.Name = buf.ReadString(int(buf.ReadUint16LE()))
	s.Id = buf.ReadInt32LE()
	s.Email = buf.ReadString(int(buf.ReadUint16LE()))
	n = int(buf.ReadUint16LE())
	s.Phone = make([]PhoneNum, n)
	for i := 0; i < n; i++ {
		(s.Phone[i]).UnmarshalBuffer(buf)
	}
}
