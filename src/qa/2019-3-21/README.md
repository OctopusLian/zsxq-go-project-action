[golang中defer的使用规则](https://studygolang.com/articles/10167)看了这篇文章， 对着最后那个代码做了实验，觉得怪怪的，输出不对，难道是与文章作者的go版本不一致，  

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(c())
}
func c() (i int) {
	defer func() { i++ }()
	return 1
}
```

这段代码输出什么，请大伙指教？