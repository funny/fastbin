package main
import "github.com/funny/binary"
func (s *SizedArray) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *SizedArray) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}
func (s *SizedArray) BinarySize() (n int) {
	for i := 0; i < 10; i++ {
		n += 1
	}
	for i := 0; i < 10; i++ {
		n += 1
	}
	for i := 0; i < 10; i++ {
		n += 1
	}
	for i := 0; i < 10; i++ {
		n += 2
	}
	for i := 0; i < 10; i++ {
		n += 2
	}
	for i := 0; i < 10; i++ {
		n += 4
	}
	for i := 0; i < 10; i++ {
		n += 4
	}
	for i := 0; i < 10; i++ {
		n += 8
	}
	for i := 0; i < 10; i++ {
		n += 8
	}
	for i := 0; i < 10; i++ {
		n += 8
	}
	for i := 0; i < 10; i++ {
		n += 8
	}
	for i := 0; i < 10; i++ {
		n += 2
		n += len((s.Field11[i]))
	}
	for i := 0; i < 10; i++ {
		n += 4
	}
	for i := 0; i < 10; i++ {
		n += 8
	}
	for i := 0; i < 10; i++ {
		n += (s.Field14[i]).BinarySize()
	}
	for i := 0; i < 10; i++ {
		n += (s.Field15[i]).BinarySize()
	}
	return
}
func (s *SizedArray) MarshalBuffer(buf *binary.Buffer) {
	for i := 0; i < 10; i++ {
		if s.Field0[i] {
			buf.WriteUint8(1)
		} else {
			buf.WriteUint8(0)
		}
	}
	for i := 0; i < 10; i++ {
		buf.WriteInt8((s.Field1[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteUint8((s.Field2[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteInt16LE((s.Field3[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteUint16LE((s.Field4[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteInt32LE((s.Field5[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteUint32LE((s.Field6[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteInt64LE((s.Field7[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteUint64LE((s.Field8[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteIntLE((s.Field9[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteUintLE((s.Field10[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteUint16LE(uint16(len((s.Field11[i]))))
		buf.WriteString((s.Field11[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteFloat32LE((s.Field12[i]))
	}
	for i := 0; i < 10; i++ {
		buf.WriteFloat64LE((s.Field13[i]))
	}
	for i := 0; i < 10; i++ {
		(s.Field14[i]).MarshalBuffer(buf)
	}
	for i := 0; i < 10; i++ {
		(s.Field15[i]).MarshalBuffer(buf)
	}
}
func (s *SizedArray) UnmarshalBuffer(buf *binary.Buffer) {
	for i := 0; i < 10; i++ {
		(s.Field0[i]) = buf.ReadUint8() > 0
	}
	for i := 0; i < 10; i++ {
		(s.Field1[i]) = buf.ReadInt8()
	}
	for i := 0; i < 10; i++ {
		(s.Field2[i]) = buf.ReadUint8()
	}
	for i := 0; i < 10; i++ {
		(s.Field3[i]) = buf.ReadInt16LE()
	}
	for i := 0; i < 10; i++ {
		(s.Field4[i]) = buf.ReadUint16LE()
	}
	for i := 0; i < 10; i++ {
		(s.Field5[i]) = buf.ReadInt32LE()
	}
	for i := 0; i < 10; i++ {
		(s.Field6[i]) = buf.ReadUint32LE()
	}
	for i := 0; i < 10; i++ {
		(s.Field7[i]) = buf.ReadInt64LE()
	}
	for i := 0; i < 10; i++ {
		(s.Field8[i]) = buf.ReadUint64LE()
	}
	for i := 0; i < 10; i++ {
		(s.Field9[i]) = buf.ReadIntLE()
	}
	for i := 0; i < 10; i++ {
		(s.Field10[i]) = buf.ReadUintLE()
	}
	for i := 0; i < 10; i++ {
		(s.Field11[i]) = buf.ReadString(int(buf.ReadUint16LE()))
	}
	for i := 0; i < 10; i++ {
		(s.Field12[i]) = buf.ReadFloat32LE()
	}
	for i := 0; i < 10; i++ {
		(s.Field13[i]) = buf.ReadFloat64LE()
	}
	for i := 0; i < 10; i++ {
		(s.Field14[i]).UnmarshalBuffer(buf)
	}
	for i := 0; i < 10; i++ {
		(s.Field15[i]).UnmarshalBuffer(buf)
	}
}
func (s *ComplexCase) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *ComplexCase) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}
func (s *ComplexCase) BinarySize() (n int) {
	n += 1
	if s.PointOfPoint != nil {
		n += 1
		if (*s.PointOfPoint) != nil {
			n += 2
			n += len((*(*s.PointOfPoint)))
		}
	}
	n += 2
	for i := 0; i < len(s.ArrayOfPoint); i++ {
		n += 1
		if (s.ArrayOfPoint[i]) != nil {
			n += 8
		}
	}
	n += 1
	if s.PointOfArray != nil {
		n += 2
		for i := 0; i < len((*s.PointOfArray)); i++ {
			n += 8
		}
	}
	n += 1
	if s.PointOfArrayOfPoint != nil {
		n += 2
		for i := 0; i < len((*s.PointOfArrayOfPoint)); i++ {
			n += 1
			if ((*s.PointOfArrayOfPoint)[i]) != nil {
				n += 2
				n += len((*((*s.PointOfArrayOfPoint)[i])))
			}
		}
	}
	n += 2
	for i := 0; i < len(s.ArrayOfArray); i++ {
		n += 2
		for ii := 0; ii < len((s.ArrayOfArray[i])); ii++ {
			n += 2
			n += len(((s.ArrayOfArray[i])[ii]))
		}
	}
	n += 2
	for i := 0; i < len(s.ArrayOfSizedArray); i++ {
		for ii := 0; ii < 10; ii++ {
			n += 8
		}
	}
	for i := 0; i < 10; i++ {
		n += 2
		for ii := 0; ii < len((s.SizedArrayOfArray[i])); ii++ {
			n += 8
		}
	}
	n += 1
	if s.WTF != nil {
		n += 2
		for i := 0; i < len((*s.WTF)); i++ {
			for ii := 0; ii < 10; ii++ {
				n += 1
				if (((*s.WTF)[i])[ii]) != nil {
					n += 1
					if (*(((*s.WTF)[i])[ii])) != nil {
						for iii := 0; iii < 11; iii++ {
							n += 2
							for iiii := 0; iiii < len(((*(*(((*s.WTF)[i])[ii])))[iii])); iiii++ {
								n += 2
								n += len((((*(*(((*s.WTF)[i])[ii])))[iii])[iiii]))
							}
						}
					}
				}
			}
		}
	}
	return
}
func (s *ComplexCase) MarshalBuffer(buf *binary.Buffer) {
	if s.PointOfPoint == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		if (*s.PointOfPoint) == nil {
			buf.WriteUint8(0)
		} else {
			buf.WriteUint8(1)
			buf.WriteUint16LE(uint16(len((*(*s.PointOfPoint)))))
			buf.WriteString((*(*s.PointOfPoint)))
		}
	}
	buf.WriteUint16LE(uint16(len(s.ArrayOfPoint)))
	for i := 0; i < len(s.ArrayOfPoint); i++ {
		if (s.ArrayOfPoint[i]) == nil {
			buf.WriteUint8(0)
		} else {
			buf.WriteUint8(1)
			buf.WriteIntLE((*(s.ArrayOfPoint[i])))
		}
	}
	if s.PointOfArray == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUint16LE(uint16(len((*s.PointOfArray))))
		for i := 0; i < len((*s.PointOfArray)); i++ {
			buf.WriteIntLE(((*s.PointOfArray)[i]))
		}
	}
	if s.PointOfArrayOfPoint == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUint16LE(uint16(len((*s.PointOfArrayOfPoint))))
		for i := 0; i < len((*s.PointOfArrayOfPoint)); i++ {
			if ((*s.PointOfArrayOfPoint)[i]) == nil {
				buf.WriteUint8(0)
			} else {
				buf.WriteUint8(1)
				buf.WriteUint16LE(uint16(len((*((*s.PointOfArrayOfPoint)[i])))))
				buf.WriteString((*((*s.PointOfArrayOfPoint)[i])))
			}
		}
	}
	buf.WriteUint16LE(uint16(len(s.ArrayOfArray)))
	for i := 0; i < len(s.ArrayOfArray); i++ {
		buf.WriteUint16LE(uint16(len((s.ArrayOfArray[i]))))
		for ii := 0; ii < len((s.ArrayOfArray[i])); ii++ {
			buf.WriteUint16LE(uint16(len(((s.ArrayOfArray[i])[ii]))))
			buf.WriteString(((s.ArrayOfArray[i])[ii]))
		}
	}
	buf.WriteUint16LE(uint16(len(s.ArrayOfSizedArray)))
	for i := 0; i < len(s.ArrayOfSizedArray); i++ {
		for ii := 0; ii < 10; ii++ {
			buf.WriteIntLE(((s.ArrayOfSizedArray[i])[ii]))
		}
	}
	for i := 0; i < 10; i++ {
		buf.WriteUint16LE(uint16(len((s.SizedArrayOfArray[i]))))
		for ii := 0; ii < len((s.SizedArrayOfArray[i])); ii++ {
			buf.WriteIntLE(((s.SizedArrayOfArray[i])[ii]))
		}
	}
	if s.WTF == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUint16LE(uint16(len((*s.WTF))))
		for i := 0; i < len((*s.WTF)); i++ {
			for ii := 0; ii < 10; ii++ {
				if (((*s.WTF)[i])[ii]) == nil {
					buf.WriteUint8(0)
				} else {
					buf.WriteUint8(1)
					if (*(((*s.WTF)[i])[ii])) == nil {
						buf.WriteUint8(0)
					} else {
						buf.WriteUint8(1)
						for iii := 0; iii < 11; iii++ {
							buf.WriteUint16LE(uint16(len(((*(*(((*s.WTF)[i])[ii])))[iii]))))
							for iiii := 0; iiii < len(((*(*(((*s.WTF)[i])[ii])))[iii])); iiii++ {
								buf.WriteUint16LE(uint16(len((((*(*(((*s.WTF)[i])[ii])))[iii])[iiii]))))
								buf.WriteString((((*(*(((*s.WTF)[i])[ii])))[iii])[iiii]))
							}
						}
					}
				}
			}
		}
	}
}
func (s *ComplexCase) UnmarshalBuffer(buf *binary.Buffer) {
	var n int
	if buf.ReadUint8() == 1 {
		if buf.ReadUint8() == 1 {
			(*(*s.PointOfPoint)) = buf.ReadString(int(buf.ReadUint16LE()))
		}
	}
	n = int(buf.ReadUint16LE())
	s.ArrayOfPoint = make([]*int, n)
	for i := 0; i < n; i++ {
		if buf.ReadUint8() == 1 {
			(*(s.ArrayOfPoint[i])) = buf.ReadIntLE()
		}
	}
	if buf.ReadUint8() == 1 {
		n = int(buf.ReadUint16LE())
		(*s.PointOfArray) = make([]int, n)
		for i := 0; i < n; i++ {
			((*s.PointOfArray)[i]) = buf.ReadIntLE()
		}
	}
	if buf.ReadUint8() == 1 {
		n = int(buf.ReadUint16LE())
		(*s.PointOfArrayOfPoint) = make([]*string, n)
		for i := 0; i < n; i++ {
			if buf.ReadUint8() == 1 {
				(*((*s.PointOfArrayOfPoint)[i])) = buf.ReadString(int(buf.ReadUint16LE()))
			}
		}
	}
	n = int(buf.ReadUint16LE())
	s.ArrayOfArray = make([][]string, n)
	for i := 0; i < n; i++ {
		n = int(buf.ReadUint16LE())
		(s.ArrayOfArray[i]) = make([]string, n)
		for ii := 0; ii < n; ii++ {
			((s.ArrayOfArray[i])[ii]) = buf.ReadString(int(buf.ReadUint16LE()))
		}
	}
	n = int(buf.ReadUint16LE())
	s.ArrayOfSizedArray = make([][10]int, n)
	for i := 0; i < n; i++ {
		for ii := 0; ii < 10; ii++ {
			((s.ArrayOfSizedArray[i])[ii]) = buf.ReadIntLE()
		}
	}
	for i := 0; i < 10; i++ {
		n = int(buf.ReadUint16LE())
		(s.SizedArrayOfArray[i]) = make([]int, n)
		for ii := 0; ii < n; ii++ {
			((s.SizedArrayOfArray[i])[ii]) = buf.ReadIntLE()
		}
	}
	if buf.ReadUint8() == 1 {
		n = int(buf.ReadUint16LE())
		(*s.WTF) = make([][10]**[11][]string, n)
		for i := 0; i < n; i++ {
			for ii := 0; ii < 10; ii++ {
				if buf.ReadUint8() == 1 {
					if buf.ReadUint8() == 1 {
						for iii := 0; iii < 11; iii++ {
							n = int(buf.ReadUint16LE())
							((*(*(((*s.WTF)[i])[ii])))[iii]) = make([]string, n)
							for iiii := 0; iiii < n; iiii++ {
								(((*(*(((*s.WTF)[i])[ii])))[iii])[iiii]) = buf.ReadString(int(buf.ReadUint16LE()))
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
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *MyType1) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}
func (s *MyType1) BinarySize() (n int) {
	n += 2
	for i := 0; i < len(s.Field1); i++ {
		n += (s.Field1[i]).BinarySize()
	}
	n += 2
	for i := 0; i < len(s.Field2); i++ {
		n += 1
		if (s.Field2[i]) != nil {
			n += (*(s.Field2[i])).BinarySize()
		}
	}
	n += 1
	if s.Field3 != nil {
		n += 2
		for i := 0; i < len((*s.Field3)); i++ {
			n += ((*s.Field3)[i]).BinarySize()
		}
	}
	for i := 0; i < 10; i++ {
		n += (s.Field4[i]).BinarySize()
	}
	for i := 0; i < 11; i++ {
		n += 1
		if (s.Field5[i]) != nil {
			n += (*(s.Field5[i])).BinarySize()
		}
	}
	n += 1
	if s.Field6 != nil {
		for i := 0; i < 12; i++ {
			n += ((*s.Field6)[i]).BinarySize()
		}
	}
	n += 2
	for i := 0; i < len(s.Field7); i++ {
		for ii := 0; ii < 13; ii++ {
			n += 1
			if ((s.Field7[i])[ii]) != nil {
				n += 2
				for iii := 0; iii < len((*((s.Field7[i])[ii]))); iii++ {
					for iiii := 0; iiii < 14; iiii++ {
						n += 1
						if (((*((s.Field7[i])[ii]))[iii])[iiii]) != nil {
							n += (*(((*((s.Field7[i])[ii]))[iii])[iiii])).BinarySize()
						}
					}
				}
			}
		}
	}
	return
}
func (s *MyType1) MarshalBuffer(buf *binary.Buffer) {
	buf.WriteUint16LE(uint16(len(s.Field1)))
	for i := 0; i < len(s.Field1); i++ {
		(s.Field1[i]).MarshalBuffer(buf)
	}
	buf.WriteUint16LE(uint16(len(s.Field2)))
	for i := 0; i < len(s.Field2); i++ {
		if (s.Field2[i]) == nil {
			buf.WriteUint8(0)
		} else {
			buf.WriteUint8(1)
			(s.Field2[i]).MarshalBuffer(buf)
		}
	}
	if s.Field3 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUint16LE(uint16(len((*s.Field3))))
		for i := 0; i < len((*s.Field3)); i++ {
			((*s.Field3)[i]).MarshalBuffer(buf)
		}
	}
	for i := 0; i < 10; i++ {
		(s.Field4[i]).MarshalBuffer(buf)
	}
	for i := 0; i < 11; i++ {
		if (s.Field5[i]) == nil {
			buf.WriteUint8(0)
		} else {
			buf.WriteUint8(1)
			(s.Field5[i]).MarshalBuffer(buf)
		}
	}
	if s.Field6 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		for i := 0; i < 12; i++ {
			((*s.Field6)[i]).MarshalBuffer(buf)
		}
	}
	buf.WriteUint16LE(uint16(len(s.Field7)))
	for i := 0; i < len(s.Field7); i++ {
		for ii := 0; ii < 13; ii++ {
			if ((s.Field7[i])[ii]) == nil {
				buf.WriteUint8(0)
			} else {
				buf.WriteUint8(1)
				buf.WriteUint16LE(uint16(len((*((s.Field7[i])[ii])))))
				for iii := 0; iii < len((*((s.Field7[i])[ii]))); iii++ {
					for iiii := 0; iiii < 14; iiii++ {
						if (((*((s.Field7[i])[ii]))[iii])[iiii]) == nil {
							buf.WriteUint8(0)
						} else {
							buf.WriteUint8(1)
							(((*((s.Field7[i])[ii]))[iii])[iiii]).MarshalBuffer(buf)
						}
					}
				}
			}
		}
	}
}
func (s *MyType1) UnmarshalBuffer(buf *binary.Buffer) {
	var n int
	n = int(buf.ReadUint16LE())
	s.Field1 = make([]MyType2, n)
	for i := 0; i < n; i++ {
		(s.Field1[i]).UnmarshalBuffer(buf)
	}
	n = int(buf.ReadUint16LE())
	s.Field2 = make([]*MyType2, n)
	for i := 0; i < n; i++ {
		if buf.ReadUint8() == 1 {
			(s.Field2[i]) = new(MyType2)
			(s.Field2[i]).UnmarshalBuffer(buf)
		}
	}
	if buf.ReadUint8() == 1 {
		n = int(buf.ReadUint16LE())
		(*s.Field3) = make([]MyType2, n)
		for i := 0; i < n; i++ {
			((*s.Field3)[i]).UnmarshalBuffer(buf)
		}
	}
	for i := 0; i < 10; i++ {
		(s.Field4[i]).UnmarshalBuffer(buf)
	}
	for i := 0; i < 11; i++ {
		if buf.ReadUint8() == 1 {
			(s.Field5[i]) = new(MyType2)
			(s.Field5[i]).UnmarshalBuffer(buf)
		}
	}
	if buf.ReadUint8() == 1 {
		for i := 0; i < 12; i++ {
			((*s.Field6)[i]).UnmarshalBuffer(buf)
		}
	}
	n = int(buf.ReadUint16LE())
	s.Field7 = make([][13]*[][14]*MyType2, n)
	for i := 0; i < n; i++ {
		for ii := 0; ii < 13; ii++ {
			if buf.ReadUint8() == 1 {
				n = int(buf.ReadUint16LE())
				(*((s.Field7[i])[ii])) = make([][14]*MyType2, n)
				for iii := 0; iii < n; iii++ {
					for iiii := 0; iiii < 14; iiii++ {
						if buf.ReadUint8() == 1 {
							(((*((s.Field7[i])[ii]))[iii])[iiii]) = new(MyType2)
							(((*((s.Field7[i])[ii]))[iii])[iiii]).UnmarshalBuffer(buf)
						}
					}
				}
			}
		}
	}
}
func (s *MyType2) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *MyType2) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}
func (s *MyType2) BinarySize() (n int) {
	n += 8
	return
}
func (s *MyType2) MarshalBuffer(buf *binary.Buffer) {
	buf.WriteIntLE(s.Field1)
}
func (s *MyType2) UnmarshalBuffer(buf *binary.Buffer) {
	s.Field1 = buf.ReadIntLE()
}
func (s *SimpleTypes) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *SimpleTypes) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}
func (s *SimpleTypes) BinarySize() (n int) {
	n += 1
	n += 1
	n += 1
	n += 2
	n += 2
	n += 4
	n += 4
	n += 8
	n += 8
	n += 8
	n += 8
	n += 2
	n += len(s.Field11)
	n += 4
	n += 8
	n += s.Field14.BinarySize()
	return
}
func (s *SimpleTypes) MarshalBuffer(buf *binary.Buffer) {
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
	buf.WriteFloat32LE(s.Field12)
	buf.WriteFloat64LE(s.Field13)
	s.Field14.MarshalBuffer(buf)
}
func (s *SimpleTypes) UnmarshalBuffer(buf *binary.Buffer) {
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
	s.Field12 = buf.ReadFloat32LE()
	s.Field13 = buf.ReadFloat64LE()
	s.Field14.UnmarshalBuffer(buf)
}
func (s *Points) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *Points) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
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
		n += 2
		n += len((*s.Field11))
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
func (s *Points) MarshalBuffer(buf *binary.Buffer) {
	if s.Field0 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		if *s.Field0 {
			buf.WriteUint8(1)
		} else {
			buf.WriteUint8(0)
		}
	}
	if s.Field1 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteInt8((*s.Field1))
	}
	if s.Field2 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUint8((*s.Field2))
	}
	if s.Field3 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteInt16LE((*s.Field3))
	}
	if s.Field4 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUint16LE((*s.Field4))
	}
	if s.Field5 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteInt32LE((*s.Field5))
	}
	if s.Field6 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUint32LE((*s.Field6))
	}
	if s.Field7 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteInt64LE((*s.Field7))
	}
	if s.Field8 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUint64LE((*s.Field8))
	}
	if s.Field9 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteIntLE((*s.Field9))
	}
	if s.Field10 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUintLE((*s.Field10))
	}
	if s.Field11 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteUint16LE(uint16(len((*s.Field11))))
		buf.WriteString((*s.Field11))
	}
	if s.Field12 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteFloat32LE((*s.Field12))
	}
	if s.Field13 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		buf.WriteFloat64LE((*s.Field13))
	}
	if s.Field14 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		s.Field14.MarshalBuffer(buf)
	}
	if s.Field15 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)
		s.Field15.MarshalBuffer(buf)
	}
}
func (s *Points) UnmarshalBuffer(buf *binary.Buffer) {
	if buf.ReadUint8() == 1 {
		(*s.Field0) = buf.ReadUint8() > 0
	}
	if buf.ReadUint8() == 1 {
		(*s.Field1) = buf.ReadInt8()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field2) = buf.ReadUint8()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field3) = buf.ReadInt16LE()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field4) = buf.ReadUint16LE()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field5) = buf.ReadInt32LE()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field6) = buf.ReadUint32LE()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field7) = buf.ReadInt64LE()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field8) = buf.ReadUint64LE()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field9) = buf.ReadIntLE()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field10) = buf.ReadUintLE()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field11) = buf.ReadString(int(buf.ReadUint16LE()))
	}
	if buf.ReadUint8() == 1 {
		(*s.Field12) = buf.ReadFloat32LE()
	}
	if buf.ReadUint8() == 1 {
		(*s.Field13) = buf.ReadFloat64LE()
	}
	if buf.ReadUint8() == 1 {
		s.Field14 = new(MyType1)
		s.Field14.UnmarshalBuffer(buf)
	}
	if buf.ReadUint8() == 1 {
		s.Field15 = new(MyType1)
		s.Field15.UnmarshalBuffer(buf)
	}
}
func (s *Arrays) MarshalBinary() (data []byte, err error) {
	var buf = binary.Buffer{Data: make([]byte, s.BinarySize())}
	s.MarshalBuffer(&buf)
	data = buf.Data[:buf.WritePos]
	return
}
func (s *Arrays) UnmarshalBinary(data []byte) error {
	s.UnmarshalBuffer(&binary.Buffer{Data: data})
	return nil
}
func (s *Arrays) BinarySize() (n int) {
	n += 2
	for i := 0; i < len(s.Field0); i++ {
		n += 1
	}
	n += 2
	for i := 0; i < len(s.Field1); i++ {
		n += 1
	}
	n += 2
	for i := 0; i < len(s.Field2); i++ {
		n += 1
	}
	n += 2
	for i := 0; i < len(s.Field3); i++ {
		n += 2
	}
	n += 2
	for i := 0; i < len(s.Field4); i++ {
		n += 2
	}
	n += 2
	for i := 0; i < len(s.Field5); i++ {
		n += 4
	}
	n += 2
	for i := 0; i < len(s.Field6); i++ {
		n += 4
	}
	n += 2
	for i := 0; i < len(s.Field7); i++ {
		n += 8
	}
	n += 2
	for i := 0; i < len(s.Field8); i++ {
		n += 8
	}
	n += 2
	for i := 0; i < len(s.Field9); i++ {
		n += 8
	}
	n += 2
	for i := 0; i < len(s.Field10); i++ {
		n += 8
	}
	n += 2
	for i := 0; i < len(s.Field11); i++ {
		n += 2
		n += len((s.Field11[i]))
	}
	n += 2
	for i := 0; i < len(s.Field12); i++ {
		n += 4
	}
	n += 2
	for i := 0; i < len(s.Field13); i++ {
		n += 8
	}
	n += 2
	for i := 0; i < len(s.Field14); i++ {
		n += (s.Field14[i]).BinarySize()
	}
	n += 2
	for i := 0; i < len(s.Field15); i++ {
		n += (s.Field15[i]).BinarySize()
	}
	return
}
func (s *Arrays) MarshalBuffer(buf *binary.Buffer) {
	buf.WriteUint16LE(uint16(len(s.Field0)))
	for i := 0; i < len(s.Field0); i++ {
		if s.Field0[i] {
			buf.WriteUint8(1)
		} else {
			buf.WriteUint8(0)
		}
	}
	buf.WriteUint16LE(uint16(len(s.Field1)))
	for i := 0; i < len(s.Field1); i++ {
		buf.WriteInt8((s.Field1[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field2)))
	for i := 0; i < len(s.Field2); i++ {
		buf.WriteUint8((s.Field2[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field3)))
	for i := 0; i < len(s.Field3); i++ {
		buf.WriteInt16LE((s.Field3[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field4)))
	for i := 0; i < len(s.Field4); i++ {
		buf.WriteUint16LE((s.Field4[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field5)))
	for i := 0; i < len(s.Field5); i++ {
		buf.WriteInt32LE((s.Field5[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field6)))
	for i := 0; i < len(s.Field6); i++ {
		buf.WriteUint32LE((s.Field6[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field7)))
	for i := 0; i < len(s.Field7); i++ {
		buf.WriteInt64LE((s.Field7[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field8)))
	for i := 0; i < len(s.Field8); i++ {
		buf.WriteUint64LE((s.Field8[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field9)))
	for i := 0; i < len(s.Field9); i++ {
		buf.WriteIntLE((s.Field9[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field10)))
	for i := 0; i < len(s.Field10); i++ {
		buf.WriteUintLE((s.Field10[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field11)))
	for i := 0; i < len(s.Field11); i++ {
		buf.WriteUint16LE(uint16(len((s.Field11[i]))))
		buf.WriteString((s.Field11[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field12)))
	for i := 0; i < len(s.Field12); i++ {
		buf.WriteFloat32LE((s.Field12[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field13)))
	for i := 0; i < len(s.Field13); i++ {
		buf.WriteFloat64LE((s.Field13[i]))
	}
	buf.WriteUint16LE(uint16(len(s.Field14)))
	for i := 0; i < len(s.Field14); i++ {
		(s.Field14[i]).MarshalBuffer(buf)
	}
	buf.WriteUint16LE(uint16(len(s.Field15)))
	for i := 0; i < len(s.Field15); i++ {
		(s.Field15[i]).MarshalBuffer(buf)
	}
}
func (s *Arrays) UnmarshalBuffer(buf *binary.Buffer) {
	var n int
	n = int(buf.ReadUint16LE())
	s.Field0 = make([]bool, n)
	for i := 0; i < n; i++ {
		(s.Field0[i]) = buf.ReadUint8() > 0
	}
	n = int(buf.ReadUint16LE())
	s.Field1 = make([]int8, n)
	for i := 0; i < n; i++ {
		(s.Field1[i]) = buf.ReadInt8()
	}
	n = int(buf.ReadUint16LE())
	s.Field2 = make([]uint8, n)
	for i := 0; i < n; i++ {
		(s.Field2[i]) = buf.ReadUint8()
	}
	n = int(buf.ReadUint16LE())
	s.Field3 = make([]int16, n)
	for i := 0; i < n; i++ {
		(s.Field3[i]) = buf.ReadInt16LE()
	}
	n = int(buf.ReadUint16LE())
	s.Field4 = make([]uint16, n)
	for i := 0; i < n; i++ {
		(s.Field4[i]) = buf.ReadUint16LE()
	}
	n = int(buf.ReadUint16LE())
	s.Field5 = make([]int32, n)
	for i := 0; i < n; i++ {
		(s.Field5[i]) = buf.ReadInt32LE()
	}
	n = int(buf.ReadUint16LE())
	s.Field6 = make([]uint32, n)
	for i := 0; i < n; i++ {
		(s.Field6[i]) = buf.ReadUint32LE()
	}
	n = int(buf.ReadUint16LE())
	s.Field7 = make([]int64, n)
	for i := 0; i < n; i++ {
		(s.Field7[i]) = buf.ReadInt64LE()
	}
	n = int(buf.ReadUint16LE())
	s.Field8 = make([]uint64, n)
	for i := 0; i < n; i++ {
		(s.Field8[i]) = buf.ReadUint64LE()
	}
	n = int(buf.ReadUint16LE())
	s.Field9 = make([]int, n)
	for i := 0; i < n; i++ {
		(s.Field9[i]) = buf.ReadIntLE()
	}
	n = int(buf.ReadUint16LE())
	s.Field10 = make([]uint, n)
	for i := 0; i < n; i++ {
		(s.Field10[i]) = buf.ReadUintLE()
	}
	n = int(buf.ReadUint16LE())
	s.Field11 = make([]string, n)
	for i := 0; i < n; i++ {
		(s.Field11[i]) = buf.ReadString(int(buf.ReadUint16LE()))
	}
	n = int(buf.ReadUint16LE())
	s.Field12 = make([]float32, n)
	for i := 0; i < n; i++ {
		(s.Field12[i]) = buf.ReadFloat32LE()
	}
	n = int(buf.ReadUint16LE())
	s.Field13 = make([]float64, n)
	for i := 0; i < n; i++ {
		(s.Field13[i]) = buf.ReadFloat64LE()
	}
	n = int(buf.ReadUint16LE())
	s.Field14 = make([]MyType1, n)
	for i := 0; i < n; i++ {
		(s.Field14[i]).UnmarshalBuffer(buf)
	}
	n = int(buf.ReadUint16LE())
	s.Field15 = make([]MyType1, n)
	for i := 0; i < n; i++ {
		(s.Field15[i]).UnmarshalBuffer(buf)
	}
}
