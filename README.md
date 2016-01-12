NOTE：此工具还在持续开发中，可能会有较大改动。
介绍
====

这个小工具可以分析指定代码中的Go结构体，并生成对应二进制序列化和反序列化代码。

可以配合[`github.com/funny/link`](https://github.com/funny/link)内置的分包协议使用，也可以单独使用。

更多介绍：[http://zhuanlan.zhihu.com/idada/20410055](http://zhuanlan.zhihu.com/idada/20410055)

基本用法
=======

fastbin将自动为代码中加了`fb:message`标签的结构体类型生成实现以下接口的方法：

```go
import "encoding"
import "github.com/funny/link"
import "github.com/funny/binary"

type FastbinMarshaler interface {
	// 将结构体的内容序列化到 BinaryWriter 中
	MarshalWriter(w binary.BinaryWriter)
}

type FastbinUnmarshaler interface {
	// 从 BinaryReader 中反序列化出结构体数据
	UnmarshalReader(r binary.BinaryReader)
}

type Fastbin interface {
	// 配合link使用所需的接口
	link.FbMessage

	// 基本的序列化反序列化接口
	FastbinMarshaler
	FastbinUnmarshaler
	
	// 可对gob起到加速作用
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
```

注释标签用法示例：

```go
//fb:message
type MyMessage struct {
		Field1 int
		Field2 string
}
```

fastbin的代码分析是以包为单位的，代码生成是以文件为单位的。

示例，分析当前目录下的包：

```
fastbin
```

目前fastbin没有额外的命令行参数，但是之后的版本中可能会通过命令行参数来支持不同语言的代码的生成。

除了命令行执行之外，也可以结合`go generate`命令使用，只在需要生成代码的文件开头加上`go generate`指令：

```go
//go:generate fastbin
package demo

//fb:message
type MyMessage struct {
		Field1 int
		Field2 string
}
```

如果你的`$GOPATH/bin`不在`$PATH`环境变量里，可以用这样的格式：`//go:generate $GOPATH/bin/fastbin`。

`go generate`的注释标签每个包只需要有一个文件有就可以，如果重复加标签就会重复生成代码。

跟`go gnerate`结合有个好出，可以不用额外写脚本，只需要执行`go generate ./...`就可以遍历生成当前目录及子目录中所有加了`go generate`注释标签的代码。

这在组织Go项目构建的时候很有用，`go generate`还支持`...`参数，这个参数会让`go gnerate`遍历`$GOPATH`下所有的包。

配合link使用
===========

fastbin除了生成结构体的二进制序列化代码以外，还可以配合link使用，变成服务接口的代码生成工具。

fastbin和link配合使用时，需要先理解服务和消息的概念，上面我们知道`fb:message`这个标签用来声明消息类型。

配合link使用时，我们需要给接口消息分配一个消息类型ID用来识别消息类型，所以标签的形式会变成这样：`fb:message = 123`，等号后面的消息ID是0 - 255之间的值。

而一个通用的网络层不可能只支持255种消息类型，所以link要求以服务为单位来组织消息，因此我们需要用到一个新的标签`fb:service = n`。

标签中的n一样是0 - 255之间的一个数，因此link支持255种服务类型，每个服务中又各支持255种消息类型，这样就足够大部分项目的使用需求了。

服务由一系列的接口方法来处理不同的消息，当一个方法签名符合以下条件时，将被fastbin识别成消息处理接口：

1. 第一个参数是`*link.Session`类型
2. 第二个参数是标注了`fb:message = n`标签的有接口消息类型

举例：

```go

//fb:service = 1
type MyService struct {
}

//fb:message = 1
type Message1 struct {
	Field1 int
	Field2 int
}

func (s *MyService) HandleMessage1(session *link.Session, msg *Message1) {
	// ...
}
```

例子中的`HandleMessage1`将被识别成`Message1`的处理接口。

需要注意，一个消息类型只能由一个服务接口处理，出现重复的消息处理接口时就会报错，即便是方法名不一样或者服务类型不一样。

还需要注意，只有直接用于接收和发送的消息才需要有消息类型ID，消息内嵌套的类型是不需要指定ID的。

当一个消息类型被指定的服务端接口接管时，这个消息类型会增加一个`ServiceID()`方法，在link发送该类型消息时用来让对方知道当前收到的是哪个服务下的哪种消息类型。

fastbin将为每个标注了`fb:service`标签的类型生成`NewRequest()`方法。

`NewRequest()`方法中会根据消息的第一个字节来识别消息类型，然后实例化对应的消息对象，并返回这种消息类型对应的处理接口。

具体请参考link的文档和生成出来的代码：

* [link主页](https://github.com/funny/link)
* [demo/service.fb.go](https://github.com/funny/fastbin/blob/master/demo/service.fb.go)

协议格式
=======

fastbin的序列化逻辑很简单，按字段顺序执行序列化和反序列化，默认使用小端格式编码多字节数值。

fastbin支持以下基本类型：

| 类型 | 字节数 |
|------|------|
| `int8`, `uint8`, `byte`, `bool` | 1 |
| `int16`, `uint16` | 2 |
| `int32`, `uint32`, `float32` | 4 |
| `int`, `uint`, `int64`, `uint64`, `float64` | 8 |
| `string`, `[]byte` | 2 + N |

fastbin支持指针类型，指针比普通类型额外多一个字节用来区分空指针，指针值为0时表示空指针，空指针的后续内容长度为0：

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

在配合link使用时，link会固定用4个字节的包头来做分包和消息分类，其中前2个字节是包体长度信息，第3个字节是服务类型ID，第4个字节是消息类型ID。

所以配合link使用时如果需要发送64K以上的消息包，请在应用协议层面添加翻页或者分帧的设计。

更详细的内容请参考生成后的代码：

* [demo/demo.fb.go](https://github.com/funny/fastbin/blob/master/demo/demo.fb.go)
* [demo/types.fb.go](https://github.com/funny/fastbin/blob/master/demo/types.fb.go)
* [demo/service.fb.go](https://github.com/funny/fastbin/blob/master/demo/service.fb.go)

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
