介绍
====

NOTE：此工具还在持续开发中，可能会有较大改动。

这是一个用来生成Go结构体二进制序列化和反序列化代码的小工具，它可以生成的代码符合[`encoding.BinaryMarshaler`](https://golang.org/pkg/encoding/#BinaryMarshaler)和[`encoding.BinaryUnmarshaler`](https://golang.org/pkg/encoding/#BinaryUnmarshaler)接口标准。

另外支持更高效的序列化和反序列化方式，可以作为[`github.com/funny/link`](https://github.com/funny/link)的通讯协议代码生成工具使用。

这个生成工具将为代码中的每个结构体生成以下方法：

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

格式
====

生成出来的序列化和反序列化格式是很简单的顺序序列化，所有的多字节数值都以小端格式编码。

支持的数据类型和生成的数据格式如下：

* `int8`，`uint8`，`byte` - 1字节
* `int16`，`uint16` - 2字节
* `int32`，`uint32` - 4字节
* `int`，`uint`，`int64，`uint64` - 8字节
* `float32` - 4字节
* `float64` - 8字节
* `string` - 2字节长度信息，后续为变长的内容
* `[]byte` - 2字节长度信息，后续为变长的内容
* `[]int` ... - 2字节长度信息，后续为变长的内容
* `struct`，`*struct` - 调用对应类型的`MarshalBuffer`和`UnmarshalBuffer`
* `[]struct`，`[]*struct` - 2字节的数组长度信息，后续每个元素分别调用对应类型的`MarshalBuffer`和`UnmarshalBuffer`

