package main

 import "github.com/funny/binary"

import "github.com/funny/link"

func (s *MyService) ServiceID() byte {
	return 1
}
func (s *MyService) DecodeRequest(p []byte) func(*link.Session) {
	switch p[0] {
	case 1:
		req := new(MyMessage1)
		req.UnmarshalPacket(p)
		return func(ss *link.Session) {
			s.HandleMessage1(ss, req)
		}
	case 2:
		req := new(MyMessage2)
		req.UnmarshalPacket(p)
		return func(ss *link.Session) {
			s.HandleMessage2(ss, req)
		}
	}
	panic("*MyService: Unknow Message Type")
}
func (s *AddressBook) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *AddressBook) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *AddressBook) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *AddressBook) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
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
func (s *Arrays) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *Arrays) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *Arrays) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *Arrays) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
}
func (s *Arrays) BinarySize() (n int) {
	
	n += len(s.Field0) * 1
	n += len(s.Field1) * 1
	n += len(s.Field2) * 1
	n += len(s.Field3) * 2
	n += len(s.Field4) * 2
	n += len(s.Field5) * 4
	n += len(s.Field6) * 4
	n += len(s.Field7) * 8
	n += len(s.Field8) * 8
	n += len(s.Field9) * 8
	n += len(s.Field10) * 8
	n += 2
	for i0 := 0; i0 < len(s.Field11); i0++ {
		n += 2 + len(s.Field11[i0])
	}
	n += len(s.Field12) * 4
	n += len(s.Field13) * 8
	n += 2
	for i0 := 0; i0 < len(s.Field14); i0++ {
		n += s.Field14[i0].BinarySize()
	}
	n += 2
	for i0 := 0; i0 < len(s.Field15); i0++ {
		n += s.Field15[i0].BinarySize()
	}
	return
}
func (s *Arrays) MarshalWriter(w binary.BinaryWriter) {
	w.WriteUint16LE(uint16(len(s.Field0)))
	for i0 := 0; i0 < len(s.Field0); i0++ {
		if s.Field0[i0] {
			w.WriteUint8(1)
		} else {
			w.WriteUint8(0)
		}
	}
	w.WriteUint16LE(uint16(len(s.Field1)))
	for i0 := 0; i0 < len(s.Field1); i0++ {
		w.WriteUint8(uint8(s.Field1[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field2)))
	for i0 := 0; i0 < len(s.Field2); i0++ {
		w.WriteUint8(uint8(s.Field2[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field3)))
	for i0 := 0; i0 < len(s.Field3); i0++ {
		w.WriteUint16LE(uint16(s.Field3[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field4)))
	for i0 := 0; i0 < len(s.Field4); i0++ {
		w.WriteUint16LE(uint16(s.Field4[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field5)))
	for i0 := 0; i0 < len(s.Field5); i0++ {
		w.WriteUint32LE(uint32(s.Field5[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field6)))
	for i0 := 0; i0 < len(s.Field6); i0++ {
		w.WriteUint32LE(uint32(s.Field6[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field7)))
	for i0 := 0; i0 < len(s.Field7); i0++ {
		w.WriteUint64LE(uint64(s.Field7[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field8)))
	for i0 := 0; i0 < len(s.Field8); i0++ {
		w.WriteUint64LE(uint64(s.Field8[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field9)))
	for i0 := 0; i0 < len(s.Field9); i0++ {
		w.WriteUint64LE(uint64(s.Field9[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field10)))
	for i0 := 0; i0 < len(s.Field10); i0++ {
		w.WriteUint64LE(uint64(s.Field10[i0]))
	}
	w.WriteUint16LE(uint16(len(s.Field11)))
	for i0 := 0; i0 < len(s.Field11); i0++ {
		w.WriteUint16LE(uint16(len(s.Field11[i0])))
		w.WriteString(s.Field11[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field12)))
	for i0 := 0; i0 < len(s.Field12); i0++ {
		w.WriteFloat32LE(s.Field12[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field13)))
	for i0 := 0; i0 < len(s.Field13); i0++ {
		w.WriteFloat64LE(s.Field13[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field14)))
	for i0 := 0; i0 < len(s.Field14); i0++ {
		s.Field14[i0].MarshalWriter(w)
	}
	w.WriteUint16LE(uint16(len(s.Field15)))
	for i0 := 0; i0 < len(s.Field15); i0++ {
		s.Field15[i0].MarshalWriter(w)
	}
}
func (s *Arrays) UnmarshalReader(r binary.BinaryReader) {
	var n int
	n = int(r.ReadUint16LE())
	s.Field0 = make([]bool, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field0[i0] = bool(r.ReadUint8() > 0)
	}
	n = int(r.ReadUint16LE())
	s.Field1 = make([]int8, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field1[i0] = int8(r.ReadUint8())
	}
	n = int(r.ReadUint16LE())
	s.Field2 = make([]uint8, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field2[i0] = uint8(r.ReadUint8())
	}
	n = int(r.ReadUint16LE())
	s.Field3 = make([]int16, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field3[i0] = int16(r.ReadUint16LE())
	}
	n = int(r.ReadUint16LE())
	s.Field4 = make([]uint16, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field4[i0] = uint16(r.ReadUint16LE())
	}
	n = int(r.ReadUint16LE())
	s.Field5 = make([]int32, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field5[i0] = int32(r.ReadUint32LE())
	}
	n = int(r.ReadUint16LE())
	s.Field6 = make([]uint32, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field6[i0] = uint32(r.ReadUint32LE())
	}
	n = int(r.ReadUint16LE())
	s.Field7 = make([]int64, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field7[i0] = int64(r.ReadUint64LE())
	}
	n = int(r.ReadUint16LE())
	s.Field8 = make([]uint64, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field8[i0] = uint64(r.ReadUint64LE())
	}
	n = int(r.ReadUint16LE())
	s.Field9 = make([]int, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field9[i0] = int(r.ReadUint64LE())
	}
	n = int(r.ReadUint16LE())
	s.Field10 = make([]uint, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field10[i0] = uint(r.ReadUint64LE())
	}
	n = int(r.ReadUint16LE())
	s.Field11 = make([]string, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field11[i0] = string(r.ReadString(int(r.ReadUint16LE())))
	}
	n = int(r.ReadUint16LE())
	s.Field12 = make([]float32, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field12[i0] = float32(r.ReadFloat32LE())
	}
	n = int(r.ReadUint16LE())
	s.Field13 = make([]float64, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field13[i0] = float64(r.ReadFloat64LE())
	}
	n = int(r.ReadUint16LE())
	s.Field14 = make([]MyType1, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field14[i0].UnmarshalReader(r)
	}
	n = int(r.ReadUint16LE())
	s.Field15 = make([]MyType1, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field15[i0].UnmarshalReader(r)
	}
}
func (s *ComplexCase) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *ComplexCase) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *ComplexCase) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *ComplexCase) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
}
func (s *ComplexCase) BinarySize() (n int) {
	
	n += 1
	if s.PointOfPoint != nil {
		n += 1
		if (*s.PointOfPoint) != nil {
			n += 2 + len((*(*s.PointOfPoint)))
		}
	}
	n += 2
	for i0 := 0; i0 < len(s.ArrayOfPoint); i0++ {
		n += 1
		if s.ArrayOfPoint[i0] != nil {
			n += 8
		}
	}
	n += 1
	if s.PointOfArray != nil {
		n += len((*s.PointOfArray)) * 8
	}
	n += 1
	if s.PointOfArrayOfPoint != nil {
		n += 2
		for i0 := 0; i0 < len((*s.PointOfArrayOfPoint)); i0++ {
			n += 1
			if (*s.PointOfArrayOfPoint)[i0] != nil {
				n += 2 + len((*(*s.PointOfArrayOfPoint)[i0]))
			}
		}
	}
	n += 2
	for i0 := 0; i0 < len(s.ArrayOfArray); i0++ {
		n += 2
		for i1 := 0; i1 < len(s.ArrayOfArray[i0]); i1++ {
			n += 2 + len(s.ArrayOfArray[i0][i1])
		}
	}
	n += 2
	for i0 := 0; i0 < len(s.ArrayOfSizedArray); i0++ {
		n += len(s.ArrayOfSizedArray[i0]) * 8
	}
	for i0 := 0; i0 < 10; i0++ {
		n += len(s.SizedArrayOfArray[i0]) * 8
	}
	n += 1
	if s.WTF != nil {
		n += 2
		for i0 := 0; i0 < len((*s.WTF)); i0++ {
			for i1 := 0; i1 < 10; i1++ {
				n += 1
				if (*s.WTF)[i0][i1] != nil {
					n += 1
					if (*(*s.WTF)[i0][i1]) != nil {
						for i2 := 0; i2 < 11; i2++ {
							n += 2
							for i3 := 0; i3 < len((*(*(*s.WTF)[i0][i1]))[i2]); i3++ {
								n += 2 + len((*(*(*s.WTF)[i0][i1]))[i2][i3])
							}
						}
					}
				}
			}
		}
	}
	return
}
func (s *ComplexCase) MarshalWriter(w binary.BinaryWriter) {
	if s.PointOfPoint == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		if (*s.PointOfPoint) == nil {
			w.WriteUint8(0)
		} else {
			w.WriteUint8(1)
			w.WriteUint16LE(uint16(len((*(*s.PointOfPoint)))))
			w.WriteString((*(*s.PointOfPoint)))
		}
	}
	w.WriteUint16LE(uint16(len(s.ArrayOfPoint)))
	for i0 := 0; i0 < len(s.ArrayOfPoint); i0++ {
		if s.ArrayOfPoint[i0] == nil {
			w.WriteUint8(0)
		} else {
			w.WriteUint8(1)
			w.WriteUint64LE(uint64((*s.ArrayOfPoint[i0])))
		}
	}
	if s.PointOfArray == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint16LE(uint16(len((*s.PointOfArray))))
		for i0 := 0; i0 < len((*s.PointOfArray)); i0++ {
			w.WriteUint64LE(uint64((*s.PointOfArray)[i0]))
		}
	}
	if s.PointOfArrayOfPoint == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint16LE(uint16(len((*s.PointOfArrayOfPoint))))
		for i0 := 0; i0 < len((*s.PointOfArrayOfPoint)); i0++ {
			if (*s.PointOfArrayOfPoint)[i0] == nil {
				w.WriteUint8(0)
			} else {
				w.WriteUint8(1)
				w.WriteUint16LE(uint16(len((*(*s.PointOfArrayOfPoint)[i0]))))
				w.WriteString((*(*s.PointOfArrayOfPoint)[i0]))
			}
		}
	}
	w.WriteUint16LE(uint16(len(s.ArrayOfArray)))
	for i0 := 0; i0 < len(s.ArrayOfArray); i0++ {
		w.WriteUint16LE(uint16(len(s.ArrayOfArray[i0])))
		for i1 := 0; i1 < len(s.ArrayOfArray[i0]); i1++ {
			w.WriteUint16LE(uint16(len(s.ArrayOfArray[i0][i1])))
			w.WriteString(s.ArrayOfArray[i0][i1])
		}
	}
	w.WriteUint16LE(uint16(len(s.ArrayOfSizedArray)))
	for i0 := 0; i0 < len(s.ArrayOfSizedArray); i0++ {
		for i1 := 0; i1 < 10; i1++ {
			w.WriteUint64LE(uint64(s.ArrayOfSizedArray[i0][i1]))
		}
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint16LE(uint16(len(s.SizedArrayOfArray[i0])))
		for i1 := 0; i1 < len(s.SizedArrayOfArray[i0]); i1++ {
			w.WriteUint64LE(uint64(s.SizedArrayOfArray[i0][i1]))
		}
	}
	if s.WTF == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint16LE(uint16(len((*s.WTF))))
		for i0 := 0; i0 < len((*s.WTF)); i0++ {
			for i1 := 0; i1 < 10; i1++ {
				if (*s.WTF)[i0][i1] == nil {
					w.WriteUint8(0)
				} else {
					w.WriteUint8(1)
					if (*(*s.WTF)[i0][i1]) == nil {
						w.WriteUint8(0)
					} else {
						w.WriteUint8(1)
						for i2 := 0; i2 < 11; i2++ {
							w.WriteUint16LE(uint16(len((*(*(*s.WTF)[i0][i1]))[i2])))
							for i3 := 0; i3 < len((*(*(*s.WTF)[i0][i1]))[i2]); i3++ {
								w.WriteUint16LE(uint16(len((*(*(*s.WTF)[i0][i1]))[i2][i3])))
								w.WriteString((*(*(*s.WTF)[i0][i1]))[i2][i3])
							}
						}
					}
				}
			}
		}
	}
}
func (s *ComplexCase) UnmarshalReader(r binary.BinaryReader) {
	var n int
	if r.ReadUint8() == 1 {
		if r.ReadUint8() == 1 {
			(*(*s.PointOfPoint)) = string(r.ReadString(int(r.ReadUint16LE())))
		}
	}
	n = int(r.ReadUint16LE())
	s.ArrayOfPoint = make([]*int, n)
	for i0 := 0; i0 < n; i0++ {
		if r.ReadUint8() == 1 {
			(*s.ArrayOfPoint[i0]) = int(r.ReadUint64LE())
		}
	}
	if r.ReadUint8() == 1 {
		n = int(r.ReadUint16LE())
		(*s.PointOfArray) = make([]int, n)
		for i0 := 0; i0 < n; i0++ {
			(*s.PointOfArray)[i0] = int(r.ReadUint64LE())
		}
	}
	if r.ReadUint8() == 1 {
		n = int(r.ReadUint16LE())
		(*s.PointOfArrayOfPoint) = make([]*string, n)
		for i0 := 0; i0 < n; i0++ {
			if r.ReadUint8() == 1 {
				(*(*s.PointOfArrayOfPoint)[i0]) = string(r.ReadString(int(r.ReadUint16LE())))
			}
		}
	}
	n = int(r.ReadUint16LE())
	s.ArrayOfArray = make([][]string, n)
	for i0 := 0; i0 < n; i0++ {
		n = int(r.ReadUint16LE())
		s.ArrayOfArray[i0] = make([]string, n)
		for i1 := 0; i1 < n; i1++ {
			s.ArrayOfArray[i0][i1] = string(r.ReadString(int(r.ReadUint16LE())))
		}
	}
	n = int(r.ReadUint16LE())
	s.ArrayOfSizedArray = make([][10]int, n)
	for i0 := 0; i0 < n; i0++ {
		for i1 := 0; i1 < 10; i1++ {
			s.ArrayOfSizedArray[i0][i1] = int(r.ReadUint64LE())
		}
	}
	for i0 := 0; i0 < 10; i0++ {
		n = int(r.ReadUint16LE())
		s.SizedArrayOfArray[i0] = make([]int, n)
		for i1 := 0; i1 < n; i1++ {
			s.SizedArrayOfArray[i0][i1] = int(r.ReadUint64LE())
		}
	}
	if r.ReadUint8() == 1 {
		n = int(r.ReadUint16LE())
		(*s.WTF) = make([][10]**[11][]string, n)
		for i0 := 0; i0 < n; i0++ {
			for i1 := 0; i1 < 10; i1++ {
				if r.ReadUint8() == 1 {
					if r.ReadUint8() == 1 {
						for i2 := 0; i2 < 11; i2++ {
							n = int(r.ReadUint16LE())
							(*(*(*s.WTF)[i0][i1]))[i2] = make([]string, n)
							for i3 := 0; i3 < n; i3++ {
								(*(*(*s.WTF)[i0][i1]))[i2][i3] = string(r.ReadString(int(r.ReadUint16LE())))
							}
						}
					}
				}
			}
		}
	}
}
func (s *MyMessage1) MessageID() byte {
	return 1
}
func (s *MyMessage1) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *MyMessage1) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *MyMessage1) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *MyMessage1) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
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
func (s *MyMessage2) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *MyMessage2) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *MyMessage2) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *MyMessage2) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
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
func (s *MyType1) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *MyType1) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *MyType1) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *MyType1) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
}
func (s *MyType1) BinarySize() (n int) {
	
	n += 2
	for i0 := 0; i0 < len(s.Field1); i0++ {
		n += s.Field1[i0].BinarySize()
	}
	n += 2
	for i0 := 0; i0 < len(s.Field2); i0++ {
		n += 1
		if s.Field2[i0] != nil {
			n += (*s.Field2[i0]).BinarySize()
		}
	}
	n += 1
	if s.Field3 != nil {
		n += 2
		for i0 := 0; i0 < len((*s.Field3)); i0++ {
			n += (*s.Field3)[i0].BinarySize()
		}
	}
	for i0 := 0; i0 < 10; i0++ {
		n += s.Field4[i0].BinarySize()
	}
	for i0 := 0; i0 < 11; i0++ {
		n += 1
		if s.Field5[i0] != nil {
			n += (*s.Field5[i0]).BinarySize()
		}
	}
	n += 1
	if s.Field6 != nil {
		for i0 := 0; i0 < 12; i0++ {
			n += (*s.Field6)[i0].BinarySize()
		}
	}
	n += 2
	for i0 := 0; i0 < len(s.Field7); i0++ {
		for i1 := 0; i1 < 13; i1++ {
			n += 1
			if s.Field7[i0][i1] != nil {
				n += 2
				for i2 := 0; i2 < len((*s.Field7[i0][i1])); i2++ {
					for i3 := 0; i3 < 14; i3++ {
						n += 1
						if (*s.Field7[i0][i1])[i2][i3] != nil {
							n += (*(*s.Field7[i0][i1])[i2][i3]).BinarySize()
						}
					}
				}
			}
		}
	}
	return
}
func (s *MyType1) MarshalWriter(w binary.BinaryWriter) {
	w.WriteUint16LE(uint16(len(s.Field1)))
	for i0 := 0; i0 < len(s.Field1); i0++ {
		s.Field1[i0].MarshalWriter(w)
	}
	w.WriteUint16LE(uint16(len(s.Field2)))
	for i0 := 0; i0 < len(s.Field2); i0++ {
		if s.Field2[i0] == nil {
			w.WriteUint8(0)
		} else {
			w.WriteUint8(1)
			s.Field2[i0].MarshalWriter(w)
		}
	}
	if s.Field3 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint16LE(uint16(len((*s.Field3))))
		for i0 := 0; i0 < len((*s.Field3)); i0++ {
			(*s.Field3)[i0].MarshalWriter(w)
		}
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field4[i0].MarshalWriter(w)
	}
	for i0 := 0; i0 < 11; i0++ {
		if s.Field5[i0] == nil {
			w.WriteUint8(0)
		} else {
			w.WriteUint8(1)
			s.Field5[i0].MarshalWriter(w)
		}
	}
	if s.Field6 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		for i0 := 0; i0 < 12; i0++ {
			(*s.Field6)[i0].MarshalWriter(w)
		}
	}
	w.WriteUint16LE(uint16(len(s.Field7)))
	for i0 := 0; i0 < len(s.Field7); i0++ {
		for i1 := 0; i1 < 13; i1++ {
			if s.Field7[i0][i1] == nil {
				w.WriteUint8(0)
			} else {
				w.WriteUint8(1)
				w.WriteUint16LE(uint16(len((*s.Field7[i0][i1]))))
				for i2 := 0; i2 < len((*s.Field7[i0][i1])); i2++ {
					for i3 := 0; i3 < 14; i3++ {
						if (*s.Field7[i0][i1])[i2][i3] == nil {
							w.WriteUint8(0)
						} else {
							w.WriteUint8(1)
							(*s.Field7[i0][i1])[i2][i3].MarshalWriter(w)
						}
					}
				}
			}
		}
	}
}
func (s *MyType1) UnmarshalReader(r binary.BinaryReader) {
	var n int
	n = int(r.ReadUint16LE())
	s.Field1 = make([]MyType2, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field1[i0].UnmarshalReader(r)
	}
	n = int(r.ReadUint16LE())
	s.Field2 = make([]*MyType2, n)
	for i0 := 0; i0 < n; i0++ {
		if r.ReadUint8() == 1 {
			s.Field2[i0] = new(MyType2)
			s.Field2[i0].UnmarshalReader(r)
		}
	}
	if r.ReadUint8() == 1 {
		n = int(r.ReadUint16LE())
		(*s.Field3) = make([]MyType2, n)
		for i0 := 0; i0 < n; i0++ {
			(*s.Field3)[i0].UnmarshalReader(r)
		}
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field4[i0].UnmarshalReader(r)
	}
	for i0 := 0; i0 < 11; i0++ {
		if r.ReadUint8() == 1 {
			s.Field5[i0] = new(MyType2)
			s.Field5[i0].UnmarshalReader(r)
		}
	}
	if r.ReadUint8() == 1 {
		for i0 := 0; i0 < 12; i0++ {
			(*s.Field6)[i0].UnmarshalReader(r)
		}
	}
	n = int(r.ReadUint16LE())
	s.Field7 = make([][13]*[][14]*MyType2, n)
	for i0 := 0; i0 < n; i0++ {
		for i1 := 0; i1 < 13; i1++ {
			if r.ReadUint8() == 1 {
				n = int(r.ReadUint16LE())
				(*s.Field7[i0][i1]) = make([][14]*MyType2, n)
				for i2 := 0; i2 < n; i2++ {
					for i3 := 0; i3 < 14; i3++ {
						if r.ReadUint8() == 1 {
							(*s.Field7[i0][i1])[i2][i3] = new(MyType2)
							(*s.Field7[i0][i1])[i2][i3].UnmarshalReader(r)
						}
					}
				}
			}
		}
	}
}
func (s *MyType2) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *MyType2) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *MyType2) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *MyType2) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
}
func (s *MyType2) BinarySize() (n int) {
	n = 8 
	return
}
func (s *MyType2) MarshalWriter(w binary.BinaryWriter) {
	w.WriteUint64LE(uint64(s.Field1))
}
func (s *MyType2) UnmarshalReader(r binary.BinaryReader) {
	s.Field1 = int(r.ReadUint64LE())
}
func (s *Person) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *Person) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *Person) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *Person) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
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
	return s.UnmarshalPacket(data)
}
func (s *PhoneNum) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *PhoneNum) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
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
func (s *Points) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *Points) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *Points) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *Points) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
}
func (s *Points) BinarySize() (n int) {
	
	n += 1
	if s.Field0 != nil {
		n += 1
	}
	n += 1
	if s.Field1 != nil {
		n += 1
	}
	n += 1
	if s.Field2 != nil {
		n += 1
	}
	n += 1
	if s.Field3 != nil {
		n += 2
	}
	n += 1
	if s.Field4 != nil {
		n += 2
	}
	n += 1
	if s.Field5 != nil {
		n += 4
	}
	n += 1
	if s.Field6 != nil {
		n += 4
	}
	n += 1
	if s.Field7 != nil {
		n += 8
	}
	n += 1
	if s.Field8 != nil {
		n += 8
	}
	n += 1
	if s.Field9 != nil {
		n += 8
	}
	n += 1
	if s.Field10 != nil {
		n += 8
	}
	n += 1
	if s.Field11 != nil {
		n += 2 + len((*s.Field11))
	}
	n += 1
	if s.Field12 != nil {
		n += 4
	}
	n += 1
	if s.Field13 != nil {
		n += 8
	}
	n += 1
	if s.Field14 != nil {
		n += (*s.Field14).BinarySize()
	}
	n += 1
	if s.Field15 != nil {
		n += (*s.Field15).BinarySize()
	}
	return
}
func (s *Points) MarshalWriter(w binary.BinaryWriter) {
	if s.Field0 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		if *s.Field0 {
			w.WriteUint8(1)
		} else {
			w.WriteUint8(0)
		}
	}
	if s.Field1 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint8(uint8((*s.Field1)))
	}
	if s.Field2 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint8(uint8((*s.Field2)))
	}
	if s.Field3 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint16LE(uint16((*s.Field3)))
	}
	if s.Field4 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint16LE(uint16((*s.Field4)))
	}
	if s.Field5 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint32LE(uint32((*s.Field5)))
	}
	if s.Field6 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint32LE(uint32((*s.Field6)))
	}
	if s.Field7 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint64LE(uint64((*s.Field7)))
	}
	if s.Field8 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint64LE(uint64((*s.Field8)))
	}
	if s.Field9 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint64LE(uint64((*s.Field9)))
	}
	if s.Field10 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint64LE(uint64((*s.Field10)))
	}
	if s.Field11 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint16LE(uint16(len((*s.Field11))))
		w.WriteString((*s.Field11))
	}
	if s.Field12 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteFloat32LE((*s.Field12))
	}
	if s.Field13 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteFloat64LE((*s.Field13))
	}
	if s.Field14 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		s.Field14.MarshalWriter(w)
	}
	if s.Field15 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		s.Field15.MarshalWriter(w)
	}
}
func (s *Points) UnmarshalReader(r binary.BinaryReader) {
	if r.ReadUint8() == 1 {
		(*s.Field0) = bool(r.ReadUint8() > 0)
	}
	if r.ReadUint8() == 1 {
		(*s.Field1) = int8(r.ReadUint8())
	}
	if r.ReadUint8() == 1 {
		(*s.Field2) = uint8(r.ReadUint8())
	}
	if r.ReadUint8() == 1 {
		(*s.Field3) = int16(r.ReadUint16LE())
	}
	if r.ReadUint8() == 1 {
		(*s.Field4) = uint16(r.ReadUint16LE())
	}
	if r.ReadUint8() == 1 {
		(*s.Field5) = int32(r.ReadUint32LE())
	}
	if r.ReadUint8() == 1 {
		(*s.Field6) = uint32(r.ReadUint32LE())
	}
	if r.ReadUint8() == 1 {
		(*s.Field7) = int64(r.ReadUint64LE())
	}
	if r.ReadUint8() == 1 {
		(*s.Field8) = uint64(r.ReadUint64LE())
	}
	if r.ReadUint8() == 1 {
		(*s.Field9) = int(r.ReadUint64LE())
	}
	if r.ReadUint8() == 1 {
		(*s.Field10) = uint(r.ReadUint64LE())
	}
	if r.ReadUint8() == 1 {
		(*s.Field11) = string(r.ReadString(int(r.ReadUint16LE())))
	}
	if r.ReadUint8() == 1 {
		(*s.Field12) = float32(r.ReadFloat32LE())
	}
	if r.ReadUint8() == 1 {
		(*s.Field13) = float64(r.ReadFloat64LE())
	}
	if r.ReadUint8() == 1 {
		s.Field14 = new(MyType1)
		s.Field14.UnmarshalReader(r)
	}
	if r.ReadUint8() == 1 {
		s.Field15 = new(MyType1)
		s.Field15.UnmarshalReader(r)
	}
}
func (s *SimpleTypes) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *SimpleTypes) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *SimpleTypes) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *SimpleTypes) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
}
func (s *SimpleTypes) BinarySize() (n int) {
	n = 1 + 1 + 1 + 2 + 2 + 4 + 4 + 8 + 8 + 8 + 8 + 4 + 8 
	n += 2 + len(s.Field11)
	n += s.Field14.BinarySize()
	return
}
func (s *SimpleTypes) MarshalWriter(w binary.BinaryWriter) {
	if s.Field0 {
		w.WriteUint8(1)
	} else {
		w.WriteUint8(0)
	}
	w.WriteUint8(uint8(s.Field1))
	w.WriteUint8(uint8(s.Field2))
	w.WriteUint16LE(uint16(s.Field3))
	w.WriteUint16LE(uint16(s.Field4))
	w.WriteUint32LE(uint32(s.Field5))
	w.WriteUint32LE(uint32(s.Field6))
	w.WriteUint64LE(uint64(s.Field7))
	w.WriteUint64LE(uint64(s.Field8))
	w.WriteUint64LE(uint64(s.Field9))
	w.WriteUint64LE(uint64(s.Field10))
	w.WriteUint16LE(uint16(len(s.Field11)))
	w.WriteString(s.Field11)
	w.WriteFloat32LE(s.Field12)
	w.WriteFloat64LE(s.Field13)
	s.Field14.MarshalWriter(w)
}
func (s *SimpleTypes) UnmarshalReader(r binary.BinaryReader) {
	s.Field0 = bool(r.ReadUint8() > 0)
	s.Field1 = int8(r.ReadUint8())
	s.Field2 = uint8(r.ReadUint8())
	s.Field3 = int16(r.ReadUint16LE())
	s.Field4 = uint16(r.ReadUint16LE())
	s.Field5 = int32(r.ReadUint32LE())
	s.Field6 = uint32(r.ReadUint32LE())
	s.Field7 = int64(r.ReadUint64LE())
	s.Field8 = uint64(r.ReadUint64LE())
	s.Field9 = int(r.ReadUint64LE())
	s.Field10 = uint(r.ReadUint64LE())
	s.Field11 = string(r.ReadString(int(r.ReadUint16LE())))
	s.Field12 = float32(r.ReadFloat32LE())
	s.Field13 = float64(r.ReadFloat64LE())
	s.Field14.UnmarshalReader(r)
}
func (s *SizedArray) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	return buf.Data, nil
}
func (s *SizedArray) UnmarshalBinary(data []byte) error {
	return s.UnmarshalPacket(data)
}
func (s *SizedArray) MarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.MarshalWriter(&buf)
	return nil
}
func (s *SizedArray) UnmarshalPacket(p []byte) error {
	var buf = binary.Buffer{Data: p}
	s.UnmarshalReader(&buf)
	return nil
}
func (s *SizedArray) BinarySize() (n int) {
	n = 10*1 + 10*1 + 10*1 + 10*2 + 10*2 + 10*4 + 10*4 + 10*8 + 10*8 + 10*8 + 10*8 + 10*4 + 10*8 
	for i0 := 0; i0 < 10; i0++ {
		n += 2 + len(s.Field11[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		n += s.Field14[i0].BinarySize()
	}
	for i0 := 0; i0 < 10; i0++ {
		n += s.Field15[i0].BinarySize()
	}
	return
}
func (s *SizedArray) MarshalWriter(w binary.BinaryWriter) {
	for i0 := 0; i0 < 10; i0++ {
		if s.Field0[i0] {
			w.WriteUint8(1)
		} else {
			w.WriteUint8(0)
		}
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint8(uint8(s.Field1[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint8(uint8(s.Field2[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint16LE(uint16(s.Field3[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint16LE(uint16(s.Field4[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint32LE(uint32(s.Field5[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint32LE(uint32(s.Field6[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint64LE(uint64(s.Field7[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint64LE(uint64(s.Field8[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint64LE(uint64(s.Field9[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint64LE(uint64(s.Field10[i0]))
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint16LE(uint16(len(s.Field11[i0])))
		w.WriteString(s.Field11[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteFloat32LE(s.Field12[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteFloat64LE(s.Field13[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field14[i0].MarshalWriter(w)
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field15[i0].MarshalWriter(w)
	}
}
func (s *SizedArray) UnmarshalReader(r binary.BinaryReader) {
	for i0 := 0; i0 < 10; i0++ {
		s.Field0[i0] = bool(r.ReadUint8() > 0)
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field1[i0] = int8(r.ReadUint8())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field2[i0] = uint8(r.ReadUint8())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field3[i0] = int16(r.ReadUint16LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field4[i0] = uint16(r.ReadUint16LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field5[i0] = int32(r.ReadUint32LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field6[i0] = uint32(r.ReadUint32LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field7[i0] = int64(r.ReadUint64LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field8[i0] = uint64(r.ReadUint64LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field9[i0] = int(r.ReadUint64LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field10[i0] = uint(r.ReadUint64LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field11[i0] = string(r.ReadString(int(r.ReadUint16LE())))
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field12[i0] = float32(r.ReadFloat32LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field13[i0] = float64(r.ReadFloat64LE())
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field14[i0].UnmarshalReader(r)
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field15[i0].UnmarshalReader(r)
	}
}






