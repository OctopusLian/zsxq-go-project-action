## 1，以下代码是否能正常编译？如果不能，哪里有问题？  

```go
package main

func main() {
	Sum(1, 2)
	len("Hello")
}

func Sum(a, b int) int {
	return a + b
}
```

### 答案  

1. Go 语言规范规定：append cap complex imag len make new real unsafe.Alignof unsafe.Offsetof unsafe.Sizeof 以上这些 builtin 函数必须接收返回值，也就是不能用于表达式语句；而且函数没有此限制；  
2. 咱们思考为什么会有此限制？以下是我的思考：  
我们知道，函数有两种作用，1）通过输入、处理、得到输出（返回值）；2）函数的副作用（比如输出到控制台）。  
有些函数，既有输出，也有副作用；有些却只有其中之一。而 builtin 中那些函数，只有输出，没有副作用，唯一的作用是返回值，因此，Go 不允许不接收返回值。而自定义函数，Go 没有也没必要检查是否只有副作用，因此允许不接收返回值。  

## 2，以下代码有问题吗？为什么？  

```go
type student struct {
    Name string
    Age  int
}

func parseStudent() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    for _, stu := range stus {
        m[stu.Name] = &stu
    }
}
```

### 答案  

这里面关键在于 for range 循环。在 for 循环中，stu 这个变量，看起来似乎每次都是新定义的，但实际上，Go 使用了同一块内存来保存它，我们可以通过在循环里输出 stu 的地址来验证这一点（fmt.Printf("%p\n", &stu))；或者换一种方式，我们在 for 循环外定义 stu，即：  

```go
var stu student
for _, stu = range stus {}
```

结果是一样的。  

可见，Go 重用了 这块内存。  

这里还得注意一点，struct 是值类型，会进行值拷贝，也就是说，stu 的内存和 stus 中 3 个元素的内存都不一样。  

因此，循环中涉及到 struct 等，要特别注意此问题，最好是定义为 指针，这里也就是：stus := []*student{} 这种。  

题外话：学过 PHP 的人，应该熟悉，PHP 中 for 循环，如果里面使用了 引用，多次循环也会有类似的坑。  