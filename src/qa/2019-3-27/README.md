## 1，Q：脚本中的编译命令：  

```sh
CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o xxx -a main.go
```

编译期间，compile进程cpu占用率大于200%；  
比普通的go build 生成的二进制程序小（41M->31M）;  

中间的CGO_ENABLED=0；-ldflags "-s -w" ；两部分该怎么理解呢？  

A：CGO_ENABLED=0 表示禁用 CGO；  

-s去掉符号表（然后panic时候的stack trace就没有任何文件名/行号信息了）’  
-w去掉DWARF调试信息，得到的程序就不能用gdb调试了。  

所以文件变小了。  

## 2，哪儿有错  

```go
func fib() func (n int) int {
	res := 0
	return func (n int) int {
		if n < 2 {
			res = 1
		} else {
			res+ = f(n-2)
		}
		return res

	}

}
```

A：f 是不是漏了  

```go
func fib() func f (n int) int {
	res := 0
	return func f (n int) int {
		if n < 2 {
			res = 1
		} else {
			res+ = f(n-2)
		}
		return res

	}
}
```
