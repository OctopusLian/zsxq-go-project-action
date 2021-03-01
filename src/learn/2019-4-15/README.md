## 1，关于可寻址（可取地址）  

1. 直接值（临时值）不能取地址；  
如：&true、&"abc"、&math.Int() 等都是非法；  

2. 字符串字节元素不能取地址；  
如：s: = "Hello World"  
_ = &(s[5])  

3. map 元素不能取地址；  
如：m := map[int]int{99:1}  
_ = &(m[99])  

4. 编译器只会自动对变量取地址，而不会自动对直接值取地址；  
如：  
```go
type T struct{}
func (t *T) f() {}
func main() {
   t := T{}
   (&t).f() // ok ，和下一句等价 
   t.f()  // ok ，将自动取地址

   (&T{}).f() // ok
   // T{}.f() // error
             // 不会自动取地址
}
```

那么为什么经常见到 &T{} 这种写法？&T{} 是为了编程方便，添加的一个语法糖 ，是下面形式的缩写，而不是临时值不能取地址的一个例外。  

```go
temp := T{}
&temp
```

## 2，关于 Go 中的函数参数按值传递的问题。  

如下代码：  

```go
func main() {
	arr := []int{2, 3, 4}

	fmt.Printf("函数前：%p\n", arr)
	printSlice(arr)
}

func printSlice(arr []int) {
	fmt.Printf("函数中：%p\n", arr)
}
```

两次输出一样吗？为什么？  