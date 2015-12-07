介绍
====

这是一个用来生成Go结构体二进制序列化和反序列化代码的小工具，它可以生成的代码符合`encoding.BinaryMarshaler`和`encoding.BinaryUnmarshaler`接口标准，同时支持更高效的序列化和反序列化方式，可以配合[`github.com/funny/link`](https://github.com/funny/link)使用，作为通讯协议代码生成工具。

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

在需要生成代码的包的根路径执行`go generate ./...`即可生成所有代码，也可以单独指定需要生成的文件，例如：`go generate demo.go`。

NOTE：此项目还在持续开发中。