NOTE：此工具还在持续开发中，可能会有较大改动。
介绍
====

这个小工具可以分析指定代码中的Go结构体，并生成对应二进制序列化和反序列化代码，它可以生成的Go代码符合[`encoding.BinaryMarshaler`](https://golang.org/pkg/encoding/#BinaryMarshaler)和[`encoding.BinaryUnmarshaler`](https://golang.org/pkg/encoding/#BinaryUnmarshaler)接口要求。

另外支持更高效的序列化和反序列化方式，可配合[`github.com/funny/link`](https://github.com/funny/link)内置的分包协议使用。

并可以加入其它编程语言代码生成的支持，可用于游戏项目的服务端和客户端通讯协议解析代码生成。

更多介绍：[http://zhuanlan.zhihu.com/idada/20410055](http://zhuanlan.zhihu.com/idada/20410055)

![title](https://camo.githubusercontent.com/a81596e605df0a3dda051668475f7ce1a85dc290/687474703a2f2f66616e6172752e636f6d2f73706f6e6765626f622d73717561726570616e74732f696d6167652f32363632382d73706f6e6765626f622d73717561726570616e74732d746f2d646f2d6c6973742e6a7067)

Go代码生成
=========

fastbin将为代码中加了`fastbin:message`标签的结构体生成以下方法：

```go
import "github.com/funny/binary"

type FastBin interface {
	// 这个方法用于实现link分包协议要求的 PacketMarshaler 接口
	BinarySize() (n int)
	
	// 这个方法用于实现link分包协议要求的 PacketMarshaler 接口
	MarshalPacket(p []byte)
	
	// 这个方法用于实现link分包协议要求的 PacketUnmarshaler 接口
	UnmarshalPacket(p []byte)

	// 将结构体的内容序列化到BinaryWriter中
	MarshalWriter(w binary.BinaryWriter)

	// 从BinaryReader中反序列化出结构体数据
	UnmarshalReader(r binary.BinaryReader)
}
```

格式示例：

```go
//fastbin:message
type MyMessage struct {
		Field1 int
		Field2 string
}
```

fastbin是一个命令行程序，在命令行中使用的方式有三种：

1. 分析并生成指定文件夹的代码：
	```
	fastbin ./
	
	fastbin ./mypackage
	
	fastbin /mypackage
	```
2. 分析并生成当前文件夹及子一级文件夹的代码：
	```
	fastbin ./...
	```
3. 分析并生成当前`$GOPATH`中所有文件夹中的代码：
	```
	fastbin ...
	```

除了命令行执行之外，也可以结合`go generate`命令使用，只在需要生成代码的文件开头加上`go generate`指令：

```go
//go:generate $GOPATH/bin/fastbin
package demo

//fastbin:message
type MyMessage struct {
		Field1 int
		Field2 string
}
```

如果你的`$GOPATH/bin`在`$PATH`环境变量里，可以用更简单的格式：`//go:generate fastbin`

`go generate`同样支持`./...`和`...`的特殊用法，可以不用在所有代码上都加指令。

格式
====

基本格式：

1. 按字段顺序执行序列化和反序列化
2. 所有的多字节数值都以小端格式编码。

支持以下基本类型：

| 类型 | 字节数 |
|------|------|
| `int8`, `uint8`, `byte`, `bool` | 1 |
| `int16`, `uint16` | 2 |
| `int32`, `uint32`, `float32` | 4 |
| `int`, `uint`, `int64`, `uint64`, `float64` | 8 |
| `string`, `[]byte` | 2 + N |

支持指针，指针类型比普通类型额外多一个字节区分空指针，指针值为0时表示空指针，空指针的后续内容长度为0：

| 类型 | 字节数 |
|------|------|
| `*int8`, `*uint8`, `*byte`, `*bool` | 1 or 1 + 1 |
| `*int16`, `*uint16` | 1 or 1 + 2 |
| `*int32`, `*uint32`, `*float32` | 1 or 1 + 4 |
| `*int`, `*uint`, `*int64`, `*uint64`, `*float64` | 1 or 1 + 8 |

支持变长数组，变长数组采用2个字节存储数组元素个数：

| 类型 | 字节数 |
|------|------|
| `[]int8`, `[]uint8`, `[]byte`, `[]bool`, `string` | 2 + N |
| `[]int16`, `[]uint16` | 2 + N * 2 |
| `[]int32`, `[]uint32`, `[]float32` | 2 + N * 4 |
| `[]int64`, `[]uint64`, `[]float64` | 2 + N * 8 |

支持定长数组，定长数组顺序循环序列化，不需要额外长度信息：

| 类型 | 字节数 |
|------|------|
| `[N]int8`, `[N]uint8`, `[N]byte`, `[N]bool` | N |
| `[N]int16`, `[N]uint16` | N * 2 |
| `[N]int32`, `[N]uint32`, `[N]float32` | N * 4 |
| `[N]int64`, `[N]uint64`, `[N]float64` | N * 8 |

支持结构体嵌套和自定义类型，基本类型以为的所有其它类型都通过`MarshalBuffer``和`UnmarshalBuffer`进行序列化和反序列化：

| 类型 | 字节数 |
|------|------|
| `MyType` | MyType.BinarySize() |
| `*MyType` | 1 or 1 + MyType.BinarySize() |
| `[]MyType` | 2 + sum(MyType.BinarySize()) |
| `[N]MyType` | sum(MyType.BinarySize()) |

支持多维数组等复杂数据结构：

| 类型 | 说明 |
|------|-----|
| `[][]int` | 二维数组 |
| `[10][]*int` | 第一唯定长的二维数组 |
| `**int` | 指向指针的指针 |
| `*[][]int` | 指向二维数组的指针 |
| `*[10]*[]**int` | 指向定长的指针的指针的数组的指针的数组的指针 |

更详细的内容请参考生成后的代码：

* [demo/demo.fast.go](https://github.com/funny/fastbin/blob/master/demo/demo.fast.go)
* [demo/types.fast.go](https://github.com/funny/fastbin/blob/master/demo/types.fast.go)

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

FAQ
===

客户端代码怎么办？
--------------

fastbin因为是给游戏项目用，所以设计时候就考虑了要支持多种语言的代码生成，结构上是比较简单容易扩展的。

添加其它语言的代码生成可以参考`golang_gen.go`和`golang_tpl.go`实现，欢迎大家提交扩展。

协议文档怎么办？
-------------

计划下个版本加入生成协议描述文档的模板，生成一份HTML文档出来，在浏览器上直接阅读。

用起来可能类似于`godoc`命令：`fastbin -S :8080`

更多的协议特性？
-------------

关于Protobuf的optional设置，可以用指针类型部分模拟，但并不完全。

fastbin的协议结构是严格的并且不向下兼容，因为我们目前项目中客户端都有热更新技术，所以这方面需求不强烈。

如果对fastbin用的二进制格式不满意，也可以替换成自己喜欢的格式，改模板就可以。

END
===

欢迎提交Issue和PR，欢迎加群讨论：188680931。
