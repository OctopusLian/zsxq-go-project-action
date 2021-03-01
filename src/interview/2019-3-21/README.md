## 1，可变参数是空接口类型  

当参数的可变参数是空接口类型时，传人空接口的切片时需要注意参数展开的问题。例如：  

```go
func main() {
    var a = []interface{}{1, 2, 3}

    fmt.Println(a)
    fmt.Println(a...)
}
```

不管是否展开，编译器都无法发现错误，但是输出是不同的。实际中可能会出现“莫名”的情况。  

## 2，recover 知识点  

以下哪些能正常捕获异常，哪些不能？  

```go
①
func main() {
    if r := recover(); r != nil {
        log.Fatal(r)
    }
    panic(123)
    if r := recover(); r != nil {
        log.Fatal(r)
    }
}
②
func main() {
    defer func() {
        if r := MyRecover(); r != nil {
            fmt.Println(r)
        }
    }()
    panic(1)
}
func MyRecover() interface{} {
    log.Println("trace...")
    return recover()
}
③
func main() {
    defer func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println(r)
            }
        }()
    }()
    panic(1)
}
④
func MyRecover() interface{} {
    return recover()
}
func main() {
    defer MyRecover()
    panic(1)
}
⑤
func main() {
    defer recover()
    panic(1)
}
⑥
func main() {
    defer func() {
        if r := recover(); r != nil { ... }
    }()
    panic(nil)
}
```

### 答案  

1. recover 必须在 defer 函数中运行；  
2. recover 必须在defer函数中直接调用才有效，也就是不能多层函数（当然，并不要求函数是匿名还是非匿名。  

所以，那道题的答案是 4、6。不过 6 中，panic 的参数，一般不应该用 nil，但不影响 recover 的使用。  

## 3，请使用 Go 实现一个函数得到两数相加结果，可用以下两种调用方式：  

sum(2,3)输出5  
sum(2)(3)输出5  
sum(2)(3)(4) 输出9  

请写出你的代码。  

这是 Go 中文网上有人发出第一真实面试题[爸爸们，救救孩子吧，面试题内心已经崩溃～](https://studygolang.com/topics/8601)  

### 答案  

主要用到了函数类型定义（并返回该函数类型）、闭包、fmt.Stringer、不定参数等知识点。  

```go
package main

import (
	"fmt"
	"strconv"
)

var total int

// 定义一个函数类型，函数的返回值是该函数类型。这个技巧可以学习一下，挺牛逼的。
// 类似的，定义结构体的时候，结构体成员可以是该结构体的指针类型。
type SumFunc func(...int) SumFunc

// SumFunc 函数类型实现 fmt.Stringer 接口。
// 这里使用这个技巧挺取巧，挺棒的。
// 根据题目的要求，一个函数似乎一会返回一个 int 类型，一会返回一个函数类型
// 这是做不到的。所以，这里的实现，永远只返回函数类型，
// 然后借助 fmt.Print 和 fmt.Stringer 接口来做到似乎间接返回了 int 类型
func (s SumFunc) String() string {
	tmpTotal := total
	total = 0
	return strconv.Itoa(tmpTotal)
}

func main() {
	
	// 这里声明和赋值分开，保证了 sum 可以在函数体中使用。
	var sum SumFunc

	sum = func(nums ...int) SumFunc {
		for _, num := range nums {
			total += num
		}
		
		return sum
	}

	fmt.Println(sum(2, 3))
	fmt.Println(sum(2)(3))
	fmt.Println(sum(2)(3)(4))
	fmt.Println(sum(2)(3)(4, 5))
}
```
