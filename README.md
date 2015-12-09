NOTE：此工具还在持续开发中，可能会有较大改动。
介绍
====

这是小工具可以分析指定代码中的Go结构体，并生成对应二进制序列化和反序列化代码，它可以生成的Go代码符合[`encoding.BinaryMarshaler`](https://golang.org/pkg/encoding/#BinaryMarshaler)和[`encoding.BinaryUnmarshaler`](https://golang.org/pkg/encoding/#BinaryUnmarshaler)接口要求。

另外支持更高效的序列化和反序列化方式，可配合[`github.com/funny/link`](https://github.com/funny/link)内置的分包协议使用。

并可以加入其它编程语言代码生成的支持，可用于游戏项目的服务端和客户端通讯协议解析代码生成。

特性
====

这个工具所使用的序列化格式并没有任何特别之处，没有Protobuf那样的可选字段，也没有Flatbuffer那样的随机访问功能，简单到可以说是简陋。之所以开发这个工具完全是从项目开发流程的角度出发。

在我们过往项目中，一直用一套自定义的协议描述语言做协议描述，然后用这些协议描述文档来生成对应的服务端和客户度代码，这也是Protobuf、Flatbuffer等工具的流程。

这种基于配置的开发流程有个问题，在开发的时候经常需要在配置文件、命令行、生成物、功能代码之间来回切换，这样的切换很容易打断功能开发的连贯性。

我希望我们的项目开发过程中，工具可以是在手边的感觉，拿来就用，用完放下，不干扰，不分散精力。这样大家可以更专注于功能本身，而不是开发流程中的某个环节。

说个关于工具和思维连贯性的笑话，有一次我要激活个软件，SN在我Gmail邮箱里，我打开Gmail邮箱时发现公司的自动科学上网打不开Gmail，于是我就开始研究公司的科学上网哪个环节出问题，最后定为到是dnsmasq的配置问题，我就去整配置。等我配置好，打开Gmail了，我对着Gmail在那里想了好久我刚才是要干什么来着？

其实这样的事情也会发生在开发过程中，我们项目开发过程中会涉及到各种工具各种配置，这些事情每一样都分散一点精力，其实无形之中给我们带来的很多损耗。这种精力损耗不但影响开发效率甚至也影响到产品质量和开发人员的积极性，因为人的精力是有限的，当一个事情做起来步骤很多很繁琐的时候，人自然而然的就会减少做的次数，甚至从潜意识上排斥去做它。

我思索下来，觉得最直接的方式就是用代码自身作为配置，开发的时候不用管配置只管开发功能，工具反过来从代码里提取信息，这样整个开发流程中就不再需要来回切换了。

以上就是这个工具存在的目的。

格式
====

采用简单的顺序序列化，所有的多字节数值都以小端格式编码。

支持的简单数据类型如下：

| 类型 | 字节数 |
|------|------|
| int8, uint8, byte, bool | 1 |
| int16, uint16 | 2 |
| int32, uint32, float32 | 4 |
| int, uint, int64, uint64, float64 | 8 |

指针类型比普通类型额外多一个字节区分空指针，指针值为0时表示空指针，空指针的后续内容长度为0：

| 类型 | 字节数 |
|------|------|
| *int8, *uint8, *byte, *bool | 1 or 1 + 1 |
| *int16, *uint16 | 1 or 1 + 2 |
| *int32, *uint32, *float32 | 1 or 1 + 4 |
| *int, *uint, *int64, *uint64, *float64 | 1 or 1 + 8 |

变长数组类型采用2个字节存储数组元素个数：

| 类型 | 字节数 |
|------|------|
| []int8, []uint8, []byte, []bool, string | 2 + N |
| []int16, []uint16 | 2 + N * 2 |
| []int32, []uint32, []float32 | 2 + N * 4 |
| []int64, []uint64, []float64 | 2 + N * 8 |

定长数组类型不需要元素个数信息：

| 类型 | 字节数 |
|------|------|
| [N]int8, [N]uint8, [N]byte, [N]bool | N |
| [N]int16, [N]uint16 | N * 2 |
| [N]int32, [N]uint32, [N]float32 | N * 4 |
| [N]int64, [N]uint64, [N]float64 | N * 8 |

所有内置类型以外的类型将通过`MarshalBuffer``和`UnmarshalBuffer`进行序列化和反序列化：

| 类型 | 字节数 |
|------|------|
| MyType | MyType.BinarySize() |
| *MyType | 1 or 1 + MyType.BinarySize() |
| []MyType | 2 + sum(MyType.BinarySize()) |
| [N]MyType | sum(MyType.BinarySize()) |

更详细的内容请参考生成后的代码：[demo/demo.fast.go](https://github.com/funny/fastbin/blob/master/demo/demo.fast.go)

关于体积和效率我按云风给sproto做的测试里的数据结构和数据做了测试。

结构如下：

```go
type AddressBook struct {
	Person []Person
}

type Person struct {
	Name  string
	Id    int32
	Email string
	Phone []PhoneNum
}

type PhoneNum struct {
	Number string
	Type   int32
}
```

测试数据如下：

```go
ab := AddressBook{[]Person{
	{"Alice", 10000, "", []PhoneNum{
		{"123456789", 1},
		{"87654321", 2},
	}},
	{"Bob", 20000, "", []PhoneNum{
		{"01234567890", 3},
	}},
}}
```

序列化后数据体积为76字节，执行1M次编码和1M次解码所需时间为：

```
Size: 76
Marshal 1M times: 125.32859ms
Unmarshal 1M times: 638.01296ms
```

反序列化过程因为有对象创建，所以开销较大，以后可以考虑加入对象池进行优化。

注：云风给sproto的测试是在lua里的，所以两者执行时间不具有可比性。

Go代码生成
=========

这个工具将为指定代码中的每个结构体生成以下方法：

```go
type FastBin interface {
	// 这个方法用于测量序列化后的数据长度
	// 用于在反序列化前一次性准备好足够大的内存空间
	// 请参考 github.com/funny/link 文档中分包协议效率的优化提示
	BinarySize() (n int)

	// 这个方法实现了 encoding.BinaryMarshaler 接口
	// 由于接口的要求是由内部返回[]byte，所以无法优化[]byte的重用
	// 建议在实际项目中避免使用
	MarshalBinary() (data []byte, err error)

	// 这个方法实现了 encoding.BinaryUnmarshaler 接口
	UnmarshalBinary(data []byte) error

	// 将结构体的内容序列化到Buffer中
	// 内部不会动态扩容，buf的内存空间必须足够长度
	MarshalBuffer(buf *binary.Buffer)

	// 从Buffer中反序列化出结构体数据
	// buff的内存空间必须足够长度
	UnmarshalBuffer(buf *binary.Buffer)
}
```

建议结合`go generate`命令使用，在需要生成代码的文件开头加上`go generate`的编译指令：

```go
//go:generate $GOPATH/bin/fastbin
package demo

type Test struct {
	Field1 int
	Field2 string
}
```

如果你的`$GOPATH/bin`在`$PATH`环境变量里，可以用更简单的指令：`//go:generate fastbin`

在需要生成代码的包的根路径执行`go generate ./...`即可生成所有代码，也可以单独指定需要生成的文件，例如：`go generate demo.go`。
