如下代码输出什么：  

```go
func main() {
    for i := 0; i < 3; i++ {
        defer func(){ println(i) } ()
    }
}

```

应该如何改进？  