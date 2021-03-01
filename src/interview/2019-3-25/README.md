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