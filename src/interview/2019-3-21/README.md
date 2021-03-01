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