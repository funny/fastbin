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
	n = 10*1 + 10*1 + 10*1 + 10*2 + 10*2 + 10*4 + 10*4 + 10*8 + 10*8 + 10*8 + 10*8 + 10*4 + 10*8 + 0

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
func (s *SizedArray) MarshalBuffer(buf *binary.Buffer) {

	for i0 := 0; i0 < 10; i0++ {

		if s.Field0[i0] {
			buf.WriteUint8(1)
		} else {
			buf.WriteUint8(0)
		}

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteInt8(s.Field1[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteUint8(s.Field2[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteInt16LE(s.Field3[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteUint16LE(s.Field4[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteInt32LE(s.Field5[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteUint32LE(s.Field6[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteInt64LE(s.Field7[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteUint64LE(s.Field8[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteIntLE(s.Field9[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteUintLE(s.Field10[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteUint16LE(uint16(len(s.Field11[i0])))
		buf.WriteString(s.Field11[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteFloat32LE(s.Field12[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteFloat64LE(s.Field13[i0])

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field14[i0].MarshalBuffer(buf)

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field15[i0].MarshalBuffer(buf)

	}

}
func (s *SizedArray) UnmarshalBuffer(buf *binary.Buffer) {

	for i0 := 0; i0 < 10; i0++ {

		s.Field0[i0] = buf.ReadUint8() > 0

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field1[i0] = buf.ReadInt8()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field2[i0] = buf.ReadUint8()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field3[i0] = buf.ReadInt16LE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field4[i0] = buf.ReadUint16LE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field5[i0] = buf.ReadInt32LE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field6[i0] = buf.ReadUint32LE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field7[i0] = buf.ReadInt64LE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field8[i0] = buf.ReadUint64LE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field9[i0] = buf.ReadIntLE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field10[i0] = buf.ReadUintLE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field11[i0] = buf.ReadString(int(buf.ReadUint16LE()))

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field12[i0] = buf.ReadFloat32LE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field13[i0] = buf.ReadFloat64LE()

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field14[i0].UnmarshalBuffer(buf)

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field15[i0].UnmarshalBuffer(buf)

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
	n = 0

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

	for i0 := 0; i0 < len(s.ArrayOfPoint); i0++ {

		if s.ArrayOfPoint[i0] == nil {
			buf.WriteUint8(0)
		} else {
			buf.WriteUint8(1)

			buf.WriteIntLE((*s.ArrayOfPoint[i0]))

		}

	}

	if s.PointOfArray == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)

		buf.WriteUint16LE(uint16(len((*s.PointOfArray))))

		for i0 := 0; i0 < len((*s.PointOfArray)); i0++ {

			buf.WriteIntLE((*s.PointOfArray)[i0])

		}

	}

	if s.PointOfArrayOfPoint == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)

		buf.WriteUint16LE(uint16(len((*s.PointOfArrayOfPoint))))

		for i0 := 0; i0 < len((*s.PointOfArrayOfPoint)); i0++ {

			if (*s.PointOfArrayOfPoint)[i0] == nil {
				buf.WriteUint8(0)
			} else {
				buf.WriteUint8(1)

				buf.WriteUint16LE(uint16(len((*(*s.PointOfArrayOfPoint)[i0]))))
				buf.WriteString((*(*s.PointOfArrayOfPoint)[i0]))

			}

		}

	}

	buf.WriteUint16LE(uint16(len(s.ArrayOfArray)))

	for i0 := 0; i0 < len(s.ArrayOfArray); i0++ {

		buf.WriteUint16LE(uint16(len(s.ArrayOfArray[i0])))

		for i1 := 0; i1 < len(s.ArrayOfArray[i0]); i1++ {

			buf.WriteUint16LE(uint16(len(s.ArrayOfArray[i0][i1])))
			buf.WriteString(s.ArrayOfArray[i0][i1])

		}

	}

	buf.WriteUint16LE(uint16(len(s.ArrayOfSizedArray)))

	for i0 := 0; i0 < len(s.ArrayOfSizedArray); i0++ {

		for i1 := 0; i1 < 10; i1++ {

			buf.WriteIntLE(s.ArrayOfSizedArray[i0][i1])

		}

	}

	for i0 := 0; i0 < 10; i0++ {

		buf.WriteUint16LE(uint16(len(s.SizedArrayOfArray[i0])))

		for i1 := 0; i1 < len(s.SizedArrayOfArray[i0]); i1++ {

			buf.WriteIntLE(s.SizedArrayOfArray[i0][i1])

		}

	}

	if s.WTF == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)

		buf.WriteUint16LE(uint16(len((*s.WTF))))

		for i0 := 0; i0 < len((*s.WTF)); i0++ {

			for i1 := 0; i1 < 10; i1++ {

				if (*s.WTF)[i0][i1] == nil {
					buf.WriteUint8(0)
				} else {
					buf.WriteUint8(1)

					if (*(*s.WTF)[i0][i1]) == nil {
						buf.WriteUint8(0)
					} else {
						buf.WriteUint8(1)

						for i2 := 0; i2 < 11; i2++ {

							buf.WriteUint16LE(uint16(len((*(*(*s.WTF)[i0][i1]))[i2])))

							for i3 := 0; i3 < len((*(*(*s.WTF)[i0][i1]))[i2]); i3++ {

								buf.WriteUint16LE(uint16(len((*(*(*s.WTF)[i0][i1]))[i2][i3])))
								buf.WriteString((*(*(*s.WTF)[i0][i1]))[i2][i3])

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

	for i0 := 0; i0 < n; i0++ {

		if buf.ReadUint8() == 1 {

			(*s.ArrayOfPoint[i0]) = buf.ReadIntLE()

		}

	}

	if buf.ReadUint8() == 1 {

		n = int(buf.ReadUint16LE())
		(*s.PointOfArray) = make([]int, n)

		for i0 := 0; i0 < n; i0++ {

			(*s.PointOfArray)[i0] = buf.ReadIntLE()

		}

	}

	if buf.ReadUint8() == 1 {

		n = int(buf.ReadUint16LE())
		(*s.PointOfArrayOfPoint) = make([]*string, n)

		for i0 := 0; i0 < n; i0++ {

			if buf.ReadUint8() == 1 {

				(*(*s.PointOfArrayOfPoint)[i0]) = buf.ReadString(int(buf.ReadUint16LE()))

			}

		}

	}

	n = int(buf.ReadUint16LE())
	s.ArrayOfArray = make([][]string, n)

	for i0 := 0; i0 < n; i0++ {

		n = int(buf.ReadUint16LE())
		s.ArrayOfArray[i0] = make([]string, n)

		for i1 := 0; i1 < n; i1++ {

			s.ArrayOfArray[i0][i1] = buf.ReadString(int(buf.ReadUint16LE()))

		}

	}

	n = int(buf.ReadUint16LE())
	s.ArrayOfSizedArray = make([][10]int, n)

	for i0 := 0; i0 < n; i0++ {

		for i1 := 0; i1 < 10; i1++ {

			s.ArrayOfSizedArray[i0][i1] = buf.ReadIntLE()

		}

	}

	for i0 := 0; i0 < 10; i0++ {

		n = int(buf.ReadUint16LE())
		s.SizedArrayOfArray[i0] = make([]int, n)

		for i1 := 0; i1 < n; i1++ {

			s.SizedArrayOfArray[i0][i1] = buf.ReadIntLE()

		}

	}

	if buf.ReadUint8() == 1 {

		n = int(buf.ReadUint16LE())
		(*s.WTF) = make([][10]**[11][]string, n)

		for i0 := 0; i0 < n; i0++ {

			for i1 := 0; i1 < 10; i1++ {

				if buf.ReadUint8() == 1 {

					if buf.ReadUint8() == 1 {

						for i2 := 0; i2 < 11; i2++ {

							n = int(buf.ReadUint16LE())
							(*(*(*s.WTF)[i0][i1]))[i2] = make([]string, n)

							for i3 := 0; i3 < n; i3++ {

								(*(*(*s.WTF)[i0][i1]))[i2][i3] = buf.ReadString(int(buf.ReadUint16LE()))

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
	n = 0

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
func (s *MyType1) MarshalBuffer(buf *binary.Buffer) {

	buf.WriteUint16LE(uint16(len(s.Field1)))

	for i0 := 0; i0 < len(s.Field1); i0++ {

		s.Field1[i0].MarshalBuffer(buf)

	}

	buf.WriteUint16LE(uint16(len(s.Field2)))

	for i0 := 0; i0 < len(s.Field2); i0++ {

		if s.Field2[i0] == nil {
			buf.WriteUint8(0)
		} else {
			buf.WriteUint8(1)

			s.Field2[i0].MarshalBuffer(buf)

		}

	}

	if s.Field3 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)

		buf.WriteUint16LE(uint16(len((*s.Field3))))

		for i0 := 0; i0 < len((*s.Field3)); i0++ {

			(*s.Field3)[i0].MarshalBuffer(buf)

		}

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field4[i0].MarshalBuffer(buf)

	}

	for i0 := 0; i0 < 11; i0++ {

		if s.Field5[i0] == nil {
			buf.WriteUint8(0)
		} else {
			buf.WriteUint8(1)

			s.Field5[i0].MarshalBuffer(buf)

		}

	}

	if s.Field6 == nil {
		buf.WriteUint8(0)
	} else {
		buf.WriteUint8(1)

		for i0 := 0; i0 < 12; i0++ {

			(*s.Field6)[i0].MarshalBuffer(buf)

		}

	}

	buf.WriteUint16LE(uint16(len(s.Field7)))

	for i0 := 0; i0 < len(s.Field7); i0++ {

		for i1 := 0; i1 < 13; i1++ {

			if s.Field7[i0][i1] == nil {
				buf.WriteUint8(0)
			} else {
				buf.WriteUint8(1)

				buf.WriteUint16LE(uint16(len((*s.Field7[i0][i1]))))

				for i2 := 0; i2 < len((*s.Field7[i0][i1])); i2++ {

					for i3 := 0; i3 < 14; i3++ {

						if (*s.Field7[i0][i1])[i2][i3] == nil {
							buf.WriteUint8(0)
						} else {
							buf.WriteUint8(1)

							(*s.Field7[i0][i1])[i2][i3].MarshalBuffer(buf)

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

	for i0 := 0; i0 < n; i0++ {

		s.Field1[i0].UnmarshalBuffer(buf)

	}

	n = int(buf.ReadUint16LE())
	s.Field2 = make([]*MyType2, n)

	for i0 := 0; i0 < n; i0++ {

		if buf.ReadUint8() == 1 {

			s.Field2[i0] = new(MyType2)
			s.Field2[i0].UnmarshalBuffer(buf)

		}

	}

	if buf.ReadUint8() == 1 {

		n = int(buf.ReadUint16LE())
		(*s.Field3) = make([]MyType2, n)

		for i0 := 0; i0 < n; i0++ {

			(*s.Field3)[i0].UnmarshalBuffer(buf)

		}

	}

	for i0 := 0; i0 < 10; i0++ {

		s.Field4[i0].UnmarshalBuffer(buf)

	}

	for i0 := 0; i0 < 11; i0++ {

		if buf.ReadUint8() == 1 {

			s.Field5[i0] = new(MyType2)
			s.Field5[i0].UnmarshalBuffer(buf)

		}

	}

	if buf.ReadUint8() == 1 {

		for i0 := 0; i0 < 12; i0++ {

			(*s.Field6)[i0].UnmarshalBuffer(buf)

		}

	}

	n = int(buf.ReadUint16LE())
	s.Field7 = make([][13]*[][14]*MyType2, n)

	for i0 := 0; i0 < n; i0++ {

		for i1 := 0; i1 < 13; i1++ {

			if buf.ReadUint8() == 1 {

				n = int(buf.ReadUint16LE())
				(*s.Field7[i0][i1]) = make([][14]*MyType2, n)

				for i2 := 0; i2 < n; i2++ {

					for i3 := 0; i3 < 14; i3++ {

						if buf.ReadUint8() == 1 {

							(*s.Field7[i0][i1])[i2][i3] = new(MyType2)
							(*s.Field7[i0][i1])[i2][i3].UnmarshalBuffer(buf)

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
	n = 8 + 0

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
	n = 1 + 1 + 1 + 2 + 2 + 4 + 4 + 8 + 8 + 8 + 8 + 4 + 8 + 0

	n += 2 + len(s.Field11)

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
	n = 0

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
	n = 0

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
func (s *Arrays) MarshalBuffer(buf *binary.Buffer) {

	buf.WriteUint16LE(uint16(len(s.Field0)))

	for i0 := 0; i0 < len(s.Field0); i0++ {

		if s.Field0[i0] {
			buf.WriteUint8(1)
		} else {
			buf.WriteUint8(0)
		}

	}

	buf.WriteUint16LE(uint16(len(s.Field1)))

	for i0 := 0; i0 < len(s.Field1); i0++ {

		buf.WriteInt8(s.Field1[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field2)))

	for i0 := 0; i0 < len(s.Field2); i0++ {

		buf.WriteUint8(s.Field2[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field3)))

	for i0 := 0; i0 < len(s.Field3); i0++ {

		buf.WriteInt16LE(s.Field3[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field4)))

	for i0 := 0; i0 < len(s.Field4); i0++ {

		buf.WriteUint16LE(s.Field4[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field5)))

	for i0 := 0; i0 < len(s.Field5); i0++ {

		buf.WriteInt32LE(s.Field5[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field6)))

	for i0 := 0; i0 < len(s.Field6); i0++ {

		buf.WriteUint32LE(s.Field6[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field7)))

	for i0 := 0; i0 < len(s.Field7); i0++ {

		buf.WriteInt64LE(s.Field7[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field8)))

	for i0 := 0; i0 < len(s.Field8); i0++ {

		buf.WriteUint64LE(s.Field8[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field9)))

	for i0 := 0; i0 < len(s.Field9); i0++ {

		buf.WriteIntLE(s.Field9[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field10)))

	for i0 := 0; i0 < len(s.Field10); i0++ {

		buf.WriteUintLE(s.Field10[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field11)))

	for i0 := 0; i0 < len(s.Field11); i0++ {

		buf.WriteUint16LE(uint16(len(s.Field11[i0])))
		buf.WriteString(s.Field11[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field12)))

	for i0 := 0; i0 < len(s.Field12); i0++ {

		buf.WriteFloat32LE(s.Field12[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field13)))

	for i0 := 0; i0 < len(s.Field13); i0++ {

		buf.WriteFloat64LE(s.Field13[i0])

	}

	buf.WriteUint16LE(uint16(len(s.Field14)))

	for i0 := 0; i0 < len(s.Field14); i0++ {

		s.Field14[i0].MarshalBuffer(buf)

	}

	buf.WriteUint16LE(uint16(len(s.Field15)))

	for i0 := 0; i0 < len(s.Field15); i0++ {

		s.Field15[i0].MarshalBuffer(buf)

	}

}
func (s *Arrays) UnmarshalBuffer(buf *binary.Buffer) {

	var n int

	n = int(buf.ReadUint16LE())
	s.Field0 = make([]bool, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field0[i0] = buf.ReadUint8() > 0

	}

	n = int(buf.ReadUint16LE())
	s.Field1 = make([]int8, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field1[i0] = buf.ReadInt8()

	}

	n = int(buf.ReadUint16LE())
	s.Field2 = make([]uint8, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field2[i0] = buf.ReadUint8()

	}

	n = int(buf.ReadUint16LE())
	s.Field3 = make([]int16, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field3[i0] = buf.ReadInt16LE()

	}

	n = int(buf.ReadUint16LE())
	s.Field4 = make([]uint16, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field4[i0] = buf.ReadUint16LE()

	}

	n = int(buf.ReadUint16LE())
	s.Field5 = make([]int32, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field5[i0] = buf.ReadInt32LE()

	}

	n = int(buf.ReadUint16LE())
	s.Field6 = make([]uint32, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field6[i0] = buf.ReadUint32LE()

	}

	n = int(buf.ReadUint16LE())
	s.Field7 = make([]int64, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field7[i0] = buf.ReadInt64LE()

	}

	n = int(buf.ReadUint16LE())
	s.Field8 = make([]uint64, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field8[i0] = buf.ReadUint64LE()

	}

	n = int(buf.ReadUint16LE())
	s.Field9 = make([]int, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field9[i0] = buf.ReadIntLE()

	}

	n = int(buf.ReadUint16LE())
	s.Field10 = make([]uint, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field10[i0] = buf.ReadUintLE()

	}

	n = int(buf.ReadUint16LE())
	s.Field11 = make([]string, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field11[i0] = buf.ReadString(int(buf.ReadUint16LE()))

	}

	n = int(buf.ReadUint16LE())
	s.Field12 = make([]float32, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field12[i0] = buf.ReadFloat32LE()

	}

	n = int(buf.ReadUint16LE())
	s.Field13 = make([]float64, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field13[i0] = buf.ReadFloat64LE()

	}

	n = int(buf.ReadUint16LE())
	s.Field14 = make([]MyType1, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field14[i0].UnmarshalBuffer(buf)

	}

	n = int(buf.ReadUint16LE())
	s.Field15 = make([]MyType1, n)

	for i0 := 0; i0 < n; i0++ {

		s.Field15[i0].UnmarshalBuffer(buf)

	}

}
