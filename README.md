介绍
====

fastbin是一个为Go结构体自动生成二进制序列化代码的工具。

fastbin的设计目的是减少DSL的使用，让开发者可以所见即所得的定义通讯协议或者二进制数据格式。

格式
====

fastbin生成的二进制序列化代码，对多字节数值统一采用小端编码，浮点数使用IEEE 754标准。

字符串和[]byte用头两个字节存储长度。

指针类型用头一个字节用来区分空指针，指针值为0时表示空指针，空指针的后续内容长度为0。

slice用头两个字节存储元素个数，之后是连续的元素数据。

数组顺序循环序列化，没有额外长度信息。

map类型用头两个字节存储元素个数，之后是连续的key-value数据。

基本类型以外的所有其它类型都通过`MarshalWriter`和`UnmarshalReader`进行序列化和反序列化。

类型
====

fastbin目前支持以下基本数据类型：

| 类型 | 字节长度 |
| --- | --- |
| int | 8 |
| uint | 8 |
| int8 | 1 |
| uint8 | 1 |
| int16 | 2 |
| uint16 | 2 |
| int32 | 4 |
| uint32 | 4 |
| int64 | 8 |
| uint64 | 8 |
| float32 | 4 |
| float64 | 8 |
| byte | 1 |
| bool | 1 |

fastbin还支持以下复杂数据类型：

| 类型 | 编码格式|
| --- | --- |
| string | 2个字节的长度信息 + 不定长的内容 |
| []byte | 2个字节的长度信息 + 不定长的内容 |
| []数组 | 2个字节的数组元素个数 + 顺序编码的数组元素 |
| [n]数组 | 顺序编码的数组元素 |
| map | 2个字节的map元素个数 + 顺序编码的key和value |
| 指针 | 一个字节的指针是否为空标志 + 顺序编码的元素（如果指针不为空）|

限制
====

fastbin对map类型有以下限制：

1. key和value均支持以下类型：

```go
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
string
```

2. 支持结构体作为key，但不支持指针作为key
3. 支持指针作为value，但不支持结构体作为value

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

fastbin目前支持两种结构体字段标签：

1. 当需要让某个字段不参与序列化时，使用`fb:"-"`标签
2. 当需要指定某个字段用具体类型参与序列化时，使用`fb:"目标类型"`标签

其中类型转换标签只支持以下类型间的互相转换：

```go
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
```

参与项目
=======

欢迎提交通过github的issues功能提交反馈或提问。

技术群：474995422