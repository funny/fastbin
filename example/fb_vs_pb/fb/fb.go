package fb

type Test struct {
	Label         string
	Type          int32
	Reps          []int64
	Optionalgroup *Test_OptionalGroup
}

type Test_OptionalGroup struct {
	RequiredField string
}
