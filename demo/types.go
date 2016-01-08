package main

// fb:message
type SimpleTypes struct {
	Field0  bool
	Field1  int8
	Field2  uint8
	Field3  int16
	Field4  uint16
	Field5  int32
	Field6  uint32
	Field7  int64
	Field8  uint64
	Field9  int
	Field10 uint
	Field11 string
	Field12 float32
	Field13 float64
	Field14 MyType1
}

// fb:message
type Points struct {
	Field0  *bool
	Field1  *int8
	Field2  *uint8
	Field3  *int16
	Field4  *uint16
	Field5  *int32
	Field6  *uint32
	Field7  *int64
	Field8  *uint64
	Field9  *int
	Field10 *uint
	Field11 *string
	Field12 *float32
	Field13 *float64
	Field14 *MyType1
	Field15 *MyType1
}

// fb:message
type Arrays struct {
	Field0  []bool
	Field1  []int8
	Field2  []uint8
	Field3  []int16
	Field4  []uint16
	Field5  []int32
	Field6  []uint32
	Field7  []int64
	Field8  []uint64
	Field9  []int
	Field10 []uint
	Field11 []string
	Field12 []float32
	Field13 []float64
	Field14 []MyType1
	Field15 []MyType1
}

// fb:message
type SizedArray struct {
	Field0  [10]bool
	Field1  [10]int8
	Field2  [10]uint8
	Field3  [10]int16
	Field4  [10]uint16
	Field5  [10]int32
	Field6  [10]uint32
	Field7  [10]int64
	Field8  [10]uint64
	Field9  [10]int
	Field10 [10]uint
	Field11 [10]string
	Field12 [10]float32
	Field13 [10]float64
	Field14 [10]MyType1
	Field15 [10]MyType1
}

// fb:message
type ComplexCase struct {
	PointOfPoint        **string
	ArrayOfPoint        []*int
	PointOfArray        *[]int
	PointOfArrayOfPoint *[]*string
	ArrayOfArray        [][]string
	ArrayOfSizedArray   [][10]int
	SizedArrayOfArray   [10][]int
	WTF                 *[][10]**[11][]string
}

// fb:message
type MyType1 struct {
	Field1 []MyType2
	Field2 []*MyType2
	Field3 *[]MyType2
	Field4 [10]MyType2
	Field5 [11]*MyType2
	Field6 *[12]MyType2
	Field7 [][13]*[][14]*MyType2
}

// fb:message
type MyType2 struct {
	Field1 int
}
