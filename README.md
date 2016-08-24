介绍
====

fastbin是一个用来生成二进制序列化和反序列化代码的工具。

fastbin的设计目的是减少DSL的使用，让开发者可以所见即所得的定义通讯协议或者二进制数据格式。

传统的协议描述语言（比如Protobuf）存在以下问题：

1. 对开发者不直观，所有功能隐藏在DSL语法背后，需要熟悉DSL语法才能使用
2. 学习DSL的过程不直观，DSL所对应的最终代码对开发者是不直接可见的，需要反复研究调试才能掌握DSL语法和对应输出的关系
3. DSL和实际业务代码是分离的，开发者在实际开发过程中需要在DSL、代码生、业务代码之间来回切换，破坏开发过程的思维连贯性

fastbin用go作为描述语言，开发者只需要定义最终在业务代码中要用到的数据结构，fastbin自动为其生成序列化和反序列化代码。

对于使用go作为第一开发语言的项目，这样的开发流程可以直接省掉DSL学习过程，以及平时开发中的DSL编写和代码生成，以及DSL输出物调试等过程，并且可以使开发者的思维更连贯。

P.S：同样的体类型信息也可以用来生成客户端的协议代码（已在实际游戏项目实践）。

类型
====

基本类型
-------

fastbin支持以下基本类型：

```go
int, uint
int8, uint8
int16, uint16
int32, uint32
int64, uint64

float32, float64

byte, bool

string

[]byte
```

复合类型
-------

fastbin支持以下复合类型：

```go
结构体：MyType
指针类型：*int, *MyType
数组类型：[10]int, [10]MyType
map类型：map[string]int, map[int]MyType
```

fastbin的map类型有以下限制：

1. key和value均支持以下类型：

```go
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
string
```

2. 支持结构体作为key，但不支持指针作为key
3. 支持指针作为value，但不支持结构体作为value

组合嵌套
----

fastbin在支持的类型范围内，支持任意程度的组合嵌套，例如以下几种情况都是支持的：

```go
**MyType

[]*MyType

[][]byte

[][]int

[][]string

map[int][]byte

map[string]*MyType

[]map[string]*MyType

map[string][]*MyType
```

标签
----

fastbin目前支持两种结构体字段标签：

1. 当需要让某个字段不参与序列化时，使用`fb:"-"`标签
2. 当需要指定某个字段用具体类型参与序列化时，使用`fb:"目标类型"`标签

其中类型转换标签只支持以下类型间的互相转换：

```go
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
```

格式
====

fastbin的序列化逻辑很简单，按字段顺序执行序列化和反序列化。

基本类型使用小端格式编码多字节数值，浮点数使用IEEE 754标准。

字符串和[]byte用头两个字节存储长度。

指针类型用头一个字节用来区分空指针，指针值为0时表示空指针，空指针的后续内容长度为0。

Slice用头两个字节存储元素个数，之后是连续的元素数据。

数组顺序循环序列化，没有额外长度信息。

map类型用头两个字节存储元素个数，之后是连续的key-value数据。

基本类型以外的所有其它类型都通过`MarshalWriter`和`UnmarshalReader`进行序列化和反序列化。

END
===

欢迎提交Issue和PR。
