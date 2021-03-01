## 1，提问：请问，从字符串里查找IP，除了正则表达式匹配，还有别的方法吗。  

进行模式匹配，正则大概是最好的方式吧！  

## 2，值方法还有什么意义？ 什么场景下会用到值方法呢？  

```go
package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p Person) SetNameV(name string) {
	p.Name = name
}

func main() {
	p := Person{Name: "lucky"}
	fmt.Println(p.Name) //lucky
	p.SetName("jucky")
	fmt.Println(p.Name) //jucky
	p.SetNameV("polaris")
	fmt.Println(p.Name) //jucky
}
```

golang 调用值方法时会对变量进行深拷贝， 修改的值并不会影响到原变量的， 那么值方法还有什么意义？ 什么场景下会用到值方法呢？  

ZCB答：我记得看一篇文章go都是值拷贝，我的理解是既然拷贝指针和拷贝值一样都是值拷贝。那就把指针理解为值的一种，只是拷贝的那个指针还指向一个值；拷贝的指针不能修改（go不能进行指针运算），但可以修改指针指向的值。  