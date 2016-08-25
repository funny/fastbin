package fb_vs_pb

import (
	"testing"

	"github.com/funny/fastbin"
	"github.com/funny/fastbin/example/fb_vs_pb/fb"
	"github.com/funny/fastbin/example/fb_vs_pb/pb"
	"github.com/golang/protobuf/proto"
)

var pbTest = &pb.Test{
	Label: proto.String("hello"),
	Type:  proto.Int32(17),
	Reps:  []int64{1, 2, 3},
	Optionalgroup: &pb.Test_OptionalGroup{
		RequiredField: proto.String("good bye"),
	},
}

var pbData, _ = proto.Marshal(pbTest)

var fbTest = &fb.Test{
	Label: "hello",
	Type:  17,
	Reps:  []int64{1, 2, 3},
	Optionalgroup: &fb.Test_OptionalGroup{
		RequiredField: "good bye",
	},
}

var fbData, _ = fbTest.MarshalBinary()

func Test_FB(t *testing.T) {
	fastbin.Register(&fb.Test{})
	fastbin.GenCode()
}

func Benchmark_PB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, _ := proto.Marshal(pbTest)
		newTest := &pb.Test{}
		proto.Unmarshal(data, newTest)
	}
}

func Benchmark_PB_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		proto.Marshal(pbTest)
	}
}

func Benchmark_PB_Unmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newTest := &pb.Test{}
		proto.Unmarshal(pbData, newTest)
	}
}

func Benchmark_FB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, _ := fbTest.MarshalBinary()
		newTest := &fb.Test{}
		newTest.UnmarshalBinary(data)
	}
}

func Benchmark_FB_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fbTest.MarshalBinary()
	}
}

func Benchmark_FB_Unmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newTest := &fb.Test{}
		newTest.UnmarshalBinary(fbData)
	}
}
