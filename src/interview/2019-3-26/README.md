## 1，以下代码是否能正常运行？为什么？  

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (*Person) Study(lang string) {
	fmt.Println("study language:", lang)
}

func main() {
	var p *Person
	p.Study("Go")
}
```

## 2，一道很基础的题，如果你不知道，需要加强基础学习哦。  

请写出以下代码的输出。  

```go
func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)  // [0,0,0,0,0,1,2,3]
}
```