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