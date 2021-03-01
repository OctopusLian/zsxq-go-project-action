# 每日一学  

## 2019/3/17  

我们知道 Go 的源文件使用 UTF-8 编码的。关于 UTF-8，如果你什么都不了解，建议查阅一些资料学习一下。今天翻阅 Go语言规范，发现编译器支持 UTF-8 带 BOM 了，说是为了兼容其他工具。这让我想起当时国人的一篇文章：[修改Go语言(golang)编译器源代码让它支持UTF-8 BOM](https://studygolang.com/articles/7273)  

## 2019/3/18  

有如下代码：  
```go
// 32 位机器
1）var x int32 = 32
2）var y int = x
3）var z rune = x
```
是否能编译通过？哪里有问题？为什么？  

- 在32位机器上，int32/ int/ rune底层都是相同的类型，即int32(int与机器字长有关,64位机器就是64位大小)。但是第二第三行都会报错,因为go语言是一种强类型语言,一般来说都是需要显式地类型转换。如int(x), rune(y)  

- 不能通过，int类型跟机器字节有关，即使在32的机器上，他们都是32个字节长度，但类型不一样，还是不能编译通过，需要显示的转换才行  

- 答案：  
Go 语言不同于大多数语言，比如 C 语言，Go 类型之前转换通常需要强制转，比如 int -> int32，即使是 32 位机器也必须显示强制转换。  
但有两个类型，byte 和 rune 比较特殊，byte 是 uint8 的别名，rune 是 int32 的别名。因此它们之间可以直接转换。  
```go
type byte = uint8
type rune = int32
```

以上是文档的定义：[Package builtin](https://docs.studygolang.com/pkg/builtin/)  
另外，关于可赋值性，Go 语言规范有明确定义。[https://docs.studygolang.com/ref/spec#Assignability](https://docs.studygolang.com/ref/spec#Assignability)  

## 2019/3/19  

有如下代码：  
```go
package main
import (
    "fmt"
)
type Personer interface{
    SayHello()
}
type Student struct{
    Name string
}
func (s *Student) SayHello(){
    fmt.Println("Hello,", s.Name)
}
func main() {
    var p Personer = Student{Name:"zhangsan"}
    p.SayHello()
} 
```
输出什么？为什么？  


