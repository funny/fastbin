package main
import "github.com/funny/binary"
func (s *Arrays) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *Arrays) UnmarshalBinary(data []byte) error {
	s.UnmarshalReader(&binary.Buffer{Data: data})
	return nil
}
func (s *Arrays) BinarySize() (n int) {
	
	n += 2
	for i0 := 0; i0 < len(s.Field0); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field1); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field2); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field3); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field4); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field5); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field6); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field7); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field8); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field9); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field10); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field11); i0++ {
		n += 2 + len(s.Field11[i0])
	}
	n += 2
	for i0 := 0; i0 < len(s.Field12); i0++ {
	}
	n += 2
	for i0 := 0; i0 < len(s.Field13); i0++ {
	}
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
		w.WriteInt8(s.Field1[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field2)))
	for i0 := 0; i0 < len(s.Field2); i0++ {
		w.WriteUint8(s.Field2[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field3)))
	for i0 := 0; i0 < len(s.Field3); i0++ {
		w.WriteInt16LE(s.Field3[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field4)))
	for i0 := 0; i0 < len(s.Field4); i0++ {
		w.WriteUint16LE(s.Field4[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field5)))
	for i0 := 0; i0 < len(s.Field5); i0++ {
		w.WriteInt32LE(s.Field5[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field6)))
	for i0 := 0; i0 < len(s.Field6); i0++ {
		w.WriteUint32LE(s.Field6[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field7)))
	for i0 := 0; i0 < len(s.Field7); i0++ {
		w.WriteInt64LE(s.Field7[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field8)))
	for i0 := 0; i0 < len(s.Field8); i0++ {
		w.WriteUint64LE(s.Field8[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field9)))
	for i0 := 0; i0 < len(s.Field9); i0++ {
		w.WriteIntLE(s.Field9[i0])
	}
	w.WriteUint16LE(uint16(len(s.Field10)))
	for i0 := 0; i0 < len(s.Field10); i0++ {
		w.WriteUintLE(s.Field10[i0])
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
		s.Field0[i0] = r.ReadUint8() > 0
	}
	n = int(r.ReadUint16LE())
	s.Field1 = make([]int8, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field1[i0] = r.ReadInt8()
	}
	n = int(r.ReadUint16LE())
	s.Field2 = make([]uint8, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field2[i0] = r.ReadUint8()
	}
	n = int(r.ReadUint16LE())
	s.Field3 = make([]int16, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field3[i0] = r.ReadInt16LE()
	}
	n = int(r.ReadUint16LE())
	s.Field4 = make([]uint16, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field4[i0] = r.ReadUint16LE()
	}
	n = int(r.ReadUint16LE())
	s.Field5 = make([]int32, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field5[i0] = r.ReadInt32LE()
	}
	n = int(r.ReadUint16LE())
	s.Field6 = make([]uint32, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field6[i0] = r.ReadUint32LE()
	}
	n = int(r.ReadUint16LE())
	s.Field7 = make([]int64, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field7[i0] = r.ReadInt64LE()
	}
	n = int(r.ReadUint16LE())
	s.Field8 = make([]uint64, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field8[i0] = r.ReadUint64LE()
	}
	n = int(r.ReadUint16LE())
	s.Field9 = make([]int, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field9[i0] = r.ReadIntLE()
	}
	n = int(r.ReadUint16LE())
	s.Field10 = make([]uint, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field10[i0] = r.ReadUintLE()
	}
	n = int(r.ReadUint16LE())
	s.Field11 = make([]string, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field11[i0] = r.ReadString(int(r.ReadUint16LE()))
	}
	n = int(r.ReadUint16LE())
	s.Field12 = make([]float32, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field12[i0] = r.ReadFloat32LE()
	}
	n = int(r.ReadUint16LE())
	s.Field13 = make([]float64, n)
	for i0 := 0; i0 < n; i0++ {
		s.Field13[i0] = r.ReadFloat64LE()
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
func (s *SizedArray) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *SizedArray) UnmarshalBinary(data []byte) error {
	s.UnmarshalReader(&binary.Buffer{Data: data})
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
		w.WriteInt8(s.Field1[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint8(s.Field2[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteInt16LE(s.Field3[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint16LE(s.Field4[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteInt32LE(s.Field5[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint32LE(s.Field6[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteInt64LE(s.Field7[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint64LE(s.Field8[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteIntLE(s.Field9[i0])
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUintLE(s.Field10[i0])
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
		s.Field0[i0] = r.ReadUint8() > 0
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field1[i0] = r.ReadInt8()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field2[i0] = r.ReadUint8()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field3[i0] = r.ReadInt16LE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field4[i0] = r.ReadUint16LE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field5[i0] = r.ReadInt32LE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field6[i0] = r.ReadUint32LE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field7[i0] = r.ReadInt64LE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field8[i0] = r.ReadUint64LE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field9[i0] = r.ReadIntLE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field10[i0] = r.ReadUintLE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field11[i0] = r.ReadString(int(r.ReadUint16LE()))
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field12[i0] = r.ReadFloat32LE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field13[i0] = r.ReadFloat64LE()
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field14[i0].UnmarshalReader(r)
	}
	for i0 := 0; i0 < 10; i0++ {
		s.Field15[i0].UnmarshalReader(r)
	}
}
func (s *ComplexCase) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *ComplexCase) UnmarshalBinary(data []byte) error {
	s.UnmarshalReader(&binary.Buffer{Data: data})
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
		}
	}
	n += 1
	if s.PointOfArray != nil {
		n += 2
		for i0 := 0; i0 < len((*s.PointOfArray)); i0++ {
		}
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
		for i1 := 0; i1 < 10; i1++ {
		}
	}
	for i0 := 0; i0 < 10; i0++ {
		n += 2
		for i1 := 0; i1 < len(s.SizedArrayOfArray[i0]); i1++ {
		}
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
			w.WriteIntLE((*s.ArrayOfPoint[i0]))
		}
	}
	if s.PointOfArray == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint16LE(uint16(len((*s.PointOfArray))))
		for i0 := 0; i0 < len((*s.PointOfArray)); i0++ {
			w.WriteIntLE((*s.PointOfArray)[i0])
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
			w.WriteIntLE(s.ArrayOfSizedArray[i0][i1])
		}
	}
	for i0 := 0; i0 < 10; i0++ {
		w.WriteUint16LE(uint16(len(s.SizedArrayOfArray[i0])))
		for i1 := 0; i1 < len(s.SizedArrayOfArray[i0]); i1++ {
			w.WriteIntLE(s.SizedArrayOfArray[i0][i1])
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
			(*(*s.PointOfPoint)) = r.ReadString(int(r.ReadUint16LE()))
		}
	}
	n = int(r.ReadUint16LE())
	s.ArrayOfPoint = make([]*int, n)
	for i0 := 0; i0 < n; i0++ {
		if r.ReadUint8() == 1 {
			(*s.ArrayOfPoint[i0]) = r.ReadIntLE()
		}
	}
	if r.ReadUint8() == 1 {
		n = int(r.ReadUint16LE())
		(*s.PointOfArray) = make([]int, n)
		for i0 := 0; i0 < n; i0++ {
			(*s.PointOfArray)[i0] = r.ReadIntLE()
		}
	}
	if r.ReadUint8() == 1 {
		n = int(r.ReadUint16LE())
		(*s.PointOfArrayOfPoint) = make([]*string, n)
		for i0 := 0; i0 < n; i0++ {
			if r.ReadUint8() == 1 {
				(*(*s.PointOfArrayOfPoint)[i0]) = r.ReadString(int(r.ReadUint16LE()))
			}
		}
	}
	n = int(r.ReadUint16LE())
	s.ArrayOfArray = make([][]string, n)
	for i0 := 0; i0 < n; i0++ {
		n = int(r.ReadUint16LE())
		s.ArrayOfArray[i0] = make([]string, n)
		for i1 := 0; i1 < n; i1++ {
			s.ArrayOfArray[i0][i1] = r.ReadString(int(r.ReadUint16LE()))
		}
	}
	n = int(r.ReadUint16LE())
	s.ArrayOfSizedArray = make([][10]int, n)
	for i0 := 0; i0 < n; i0++ {
		for i1 := 0; i1 < 10; i1++ {
			s.ArrayOfSizedArray[i0][i1] = r.ReadIntLE()
		}
	}
	for i0 := 0; i0 < 10; i0++ {
		n = int(r.ReadUint16LE())
		s.SizedArrayOfArray[i0] = make([]int, n)
		for i1 := 0; i1 < n; i1++ {
			s.SizedArrayOfArray[i0][i1] = r.ReadIntLE()
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
								(*(*(*s.WTF)[i0][i1]))[i2][i3] = r.ReadString(int(r.ReadUint16LE()))
							}
						}
					}
				}
			}
		}
	}
}
func (s *MyType1) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *MyType1) UnmarshalBinary(data []byte) error {
	s.UnmarshalReader(&binary.Buffer{Data: data})
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
	data = buf.Data[:buf.WritePos]
	return
}
func (s *MyType2) UnmarshalBinary(data []byte) error {
	s.UnmarshalReader(&binary.Buffer{Data: data})
	return nil
}
func (s *MyType2) BinarySize() (n int) {
	n = 8 
	return
}
func (s *MyType2) MarshalWriter(w binary.BinaryWriter) {
	w.WriteIntLE(s.Field1)
}
func (s *MyType2) UnmarshalReader(r binary.BinaryReader) {
	s.Field1 = r.ReadIntLE()
}
func (s *SimpleTypes) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *SimpleTypes) UnmarshalBinary(data []byte) error {
	s.UnmarshalReader(&binary.Buffer{Data: data})
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
	w.WriteInt8(s.Field1)
	w.WriteUint8(s.Field2)
	w.WriteInt16LE(s.Field3)
	w.WriteUint16LE(s.Field4)
	w.WriteInt32LE(s.Field5)
	w.WriteUint32LE(s.Field6)
	w.WriteInt64LE(s.Field7)
	w.WriteUint64LE(s.Field8)
	w.WriteIntLE(s.Field9)
	w.WriteUintLE(s.Field10)
	w.WriteUint16LE(uint16(len(s.Field11)))
	w.WriteString(s.Field11)
	w.WriteFloat32LE(s.Field12)
	w.WriteFloat64LE(s.Field13)
	s.Field14.MarshalWriter(w)
}
func (s *SimpleTypes) UnmarshalReader(r binary.BinaryReader) {
	s.Field0 = r.ReadUint8() > 0
	s.Field1 = r.ReadInt8()
	s.Field2 = r.ReadUint8()
	s.Field3 = r.ReadInt16LE()
	s.Field4 = r.ReadUint16LE()
	s.Field5 = r.ReadInt32LE()
	s.Field6 = r.ReadUint32LE()
	s.Field7 = r.ReadInt64LE()
	s.Field8 = r.ReadUint64LE()
	s.Field9 = r.ReadIntLE()
	s.Field10 = r.ReadUintLE()
	s.Field11 = r.ReadString(int(r.ReadUint16LE()))
	s.Field12 = r.ReadFloat32LE()
	s.Field13 = r.ReadFloat64LE()
	s.Field14.UnmarshalReader(r)
}
func (s *Points) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalWriter(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *Points) UnmarshalBinary(data []byte) error {
	s.UnmarshalReader(&binary.Buffer{Data: data})
	return nil
}
func (s *Points) BinarySize() (n int) {
	
	n += 1
	if s.Field0 != nil {
	}
	n += 1
	if s.Field1 != nil {
	}
	n += 1
	if s.Field2 != nil {
	}
	n += 1
	if s.Field3 != nil {
	}
	n += 1
	if s.Field4 != nil {
	}
	n += 1
	if s.Field5 != nil {
	}
	n += 1
	if s.Field6 != nil {
	}
	n += 1
	if s.Field7 != nil {
	}
	n += 1
	if s.Field8 != nil {
	}
	n += 1
	if s.Field9 != nil {
	}
	n += 1
	if s.Field10 != nil {
	}
	n += 1
	if s.Field11 != nil {
		n += 2 + len((*s.Field11))
	}
	n += 1
	if s.Field12 != nil {
	}
	n += 1
	if s.Field13 != nil {
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
		w.WriteInt8((*s.Field1))
	}
	if s.Field2 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint8((*s.Field2))
	}
	if s.Field3 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteInt16LE((*s.Field3))
	}
	if s.Field4 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint16LE((*s.Field4))
	}
	if s.Field5 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteInt32LE((*s.Field5))
	}
	if s.Field6 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint32LE((*s.Field6))
	}
	if s.Field7 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteInt64LE((*s.Field7))
	}
	if s.Field8 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUint64LE((*s.Field8))
	}
	if s.Field9 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteIntLE((*s.Field9))
	}
	if s.Field10 == nil {
		w.WriteUint8(0)
	} else {
		w.WriteUint8(1)
		w.WriteUintLE((*s.Field10))
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
		(*s.Field0) = r.ReadUint8() > 0
	}
	if r.ReadUint8() == 1 {
		(*s.Field1) = r.ReadInt8()
	}
	if r.ReadUint8() == 1 {
		(*s.Field2) = r.ReadUint8()
	}
	if r.ReadUint8() == 1 {
		(*s.Field3) = r.ReadInt16LE()
	}
	if r.ReadUint8() == 1 {
		(*s.Field4) = r.ReadUint16LE()
	}
	if r.ReadUint8() == 1 {
		(*s.Field5) = r.ReadInt32LE()
	}
	if r.ReadUint8() == 1 {
		(*s.Field6) = r.ReadUint32LE()
	}
	if r.ReadUint8() == 1 {
		(*s.Field7) = r.ReadInt64LE()
	}
	if r.ReadUint8() == 1 {
		(*s.Field8) = r.ReadUint64LE()
	}
	if r.ReadUint8() == 1 {
		(*s.Field9) = r.ReadIntLE()
	}
	if r.ReadUint8() == 1 {
		(*s.Field10) = r.ReadUintLE()
	}
	if r.ReadUint8() == 1 {
		(*s.Field11) = r.ReadString(int(r.ReadUint16LE()))
	}
	if r.ReadUint8() == 1 {
		(*s.Field12) = r.ReadFloat32LE()
	}
	if r.ReadUint8() == 1 {
		(*s.Field13) = r.ReadFloat64LE()
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
