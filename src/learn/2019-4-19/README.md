## 1，关于类型  

Go 中的类型可以分为命名类型（named type） 和未命名类型（unnamed type）。命名类型包括 bool、int、string 等，⽽而 array、slice、map 等和具体元素类型、长度等有关，属于未命名类型。  

具有相同声明的未命名类型被视为同一类型。  

• 具有相同基类型的指针。  
• 具有相同元素类型和⻓度的 array。  
• 具有相同元素类型的 slice。  
• 具有相同键值类型的 map。  
• 具有相同元素类型和传送⽅方向的 channel。  
• 具有相同字段序列 (字段名、类型、标签、顺序) 的匿名 struct。  
• 签名相同 (参数和返回值，不包括参数名称) 的 function。  
• ⽅法集相同 (方法名、方法签名相同，和次序⽆无关) 的 interface。  

巩固一下：  

1.  type MyMap map[string]string  

这是命名类型还是未命名类型？  

## 2，2. 以下代码是否有问题，哪里有问题，为什么？  

```go
package main

import (
	"fmt"
)

func main() {
	type MyMap1 map[string]string
	
	type MyMap2 map[string]string
	
	var myMap = map[string]string{"name": "polaris"}
	var myMap1 MyMap1 = myMap
	var myMap2 MyMap2 = myMap1
	
	fmt.Println(myMap2)
}
```
