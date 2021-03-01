## 1,为什么 T 和 *T 有不同的方法集  

如果一个接口值包含一个指针 *T，一个方法调用可通过解引用该指针来获得一个值，所以，*T 方法集包含 T 的方法集；  

但反过来，如果一个接口值包含一个值 T，就没有安全的方式让一个方法调用获得一个指针。一方面，有可能 T 不可寻址；另一方面，即使 T 可寻址，但方法可能错误的通过指针修改它的值，而实际上这个修改会丢失，得到不是期望的结果。因此Go语言规范才定义 T 的方法集不包括 *T 的方法集。  

之前有一个主题讲解了为什么 T 的方法集不包含 *T 的方法集，但要注意，方法集的概念是用来判断类型是否实现了某个接口，不能因为 T 的方法集没有包含 *T 的方法集，就以为 T 就不能调用 *T 的方法，Go 语言规范明确说了，直接调用指针的方法，编译器会自动取 T 的指针，然后调用，这跟方法集没有关系。看如下例子，注意注释部分。希望大家明白这两者的区别和使用场景。  

```go
package main

import (
	"fmt"
)

type Speaker interface {
	Speak(language string) 
}

type Chinese struct {
	Name string
}

func (c Chinese) Speak(language string) {
	fmt.Println("My name is", c.Name, ", I am Chinese, I speak", language)
}

type American struct {
	Name string
}

func (a *American) Speak(language string) {
	fmt.Println("My name is", a.Name, ", I am American, I speak", language)
}

func main() {
	var speaker Speaker = Chinese{Name: "zhangsan"}
	speaker.Speak("Chinese")
	
	// &Chinese{} 返回的指针，虽然 Chinese 的 Speek() 方法接收者是值类型，它会包含在指针类型 *Chinese 中，因此 *Chinese 实现了 Speaker 接口
	speaker = &Chinese{Name: "lisi"}
	speaker.Speak("English")
	
	// 编译不通过， American{}返回的是值，方法集中没有 Speak() 方法，因此没有实现 Speaker 接口
	// speaker = American{Name: "John"}
	// speaker.Speak("Chinese")
	// 但是 American 的值类型却可以直接调用指针接收者的方法，如：
	american := American{Name: "Tom"}
	american.Speak("English")
	
	speaker = &American{Name: "Mary"}
	speaker.Speak("English")
}
```