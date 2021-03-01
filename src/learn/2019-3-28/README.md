## 1，关于接口的 nil 问题。  

在底层，接口作为两个元素实现：一个类型和一个值。该值被称为接口的动态值， 它是一个任意的具体值，而该接口的类型则为该值的类型。对于 int 值3， 一个接口值示意性地包含(int, 3)。  

只有在内部值和类型都未设置时(nil, nil)，一个接口的值才为 nil。特别是，一个 nil 接口将总是拥有一个 nil 类型。若我们在一个接口值中存储一个 *int 类型的指针，则内部类型将为 *int，无论该指针的值是什么：(*int, nil)。 因此，这样的接口值会是非 nil 的，即使在该指针的内部为 nil。  

这种情况会让人迷惑，而且当 nil 值存储在接口值内部时这种情况总是发生， 例如错误返回：  

```go
func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = ErrBad
	}
	return p // 将总是返回一个非nil错误。
}
```

如果一切顺利，该函数会返回一个 nil 的 p， 因此该返回值为拥有(*MyError, nil)的 error 接口值。这也就意味着如果调用者将返回的错误与 nil 相比较， 它将总是看上去有错误，即便没有什么坏事发生。要向调用者返回一个适当的 nil error，该函数必须返回一个显式的 nil：  

```go
func returnsError() error {
	if bad() {
		return ErrBad
	}
	return nil
}
```

这对于总是在签名中使用 error 类型返回错误（正如我们上面做的）而非像 *MyError 这样具体类型的函数来说是个不错的主意，它可以帮助确保错误被正确地创建。 例如，即使 os.Open 返回一个 error， 若非 nil 的话，它总是具体的类型 *os.PathError。  

对于那些描述，无论接口是否被使用，相似的情形都会出现。只要记住，如果任何具体的值已被存储在接口中， 该接口就不为 nil。  

接口的内部定义  

```go
type iface struct {
	tab  *itab // 代表类型
	data unsafe.Pointer  // 代表数据
}
```

## 2,关于类型断言  

Go 语言规范规定，类型断言是指：  

对于接口类型的表达式 x 与类型 T，主表达式：x.(T)，断言 x 不为 nil 且存储于 x 中的值其类型为 T。 记法 x.(T) 称为 类型断言。  

这里明确指出，进行类型断言时，x 必须是接口（注意，只要是接口就可以，不在乎是不是空接口）。注意，T 可以是类型或接口。  

注意，如果类型断言成立，则该表达式的值即为存储于 x 中的值，且其类型为 T；若该类型断言不成立， 就会 panic。  

```go
var x interface{} = 7  // x 拥有动态类型 int 与值 7
i := x.(int)           // i 拥有类型 int 与值 7

type I interface { m() }
var y I
s := y.(string)        // 非法：string 没有实现 I（缺少方法 m）
r := y.(io.Reader)     // r 拥有 类型 io.Reader 且 y 必须同时实现了 I 和 io.Reader，否则 panic
```

如果不确定接口的动态类型是什么，为了避免 panic，可以接收表达式的第2个参数，即：  

```go
v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
```

通过判断 ok 是 true 还是 false，如果为 true，表示 x 的动态类型是 T；否则 x 的动态类型不是 T。  

另外，x.(T) 语法中，如果 T 是接口，编译器会自动检测 x 的动态类型是否实现了接口 T。  

```go
type Shape interface {
	Area() float32
}

type Perimeter interface {
	P() float32
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * (c.radius * c.radius)
}

func main() {
	var s Shape = Circle{3}
	v1,ok1 := s.(Shape)
	v2,ok2 := s.(Perimeter)
	fmt.Println(v1,ok1)
	fmt.Println(v2,ok2)
}
```

## 3，关于类型选择 (type switch）  

类型选择的语法和类型断言的语法类似，但有如下要求：  

1. 只能用于 switch 语句；  
2. x.(type) 中的 type 是固定的，只能是 type 这个关键词；  
3. 不允许使用 fallthrough 语句；  

例如：  

```go
switch x.(type) {
// case
}
```

和类型断言一样，x 必须是接口。  

每一个 case 中的类型必须实现了 x 接口。  

另外一种语法：  

```go
switch i := x.(type) {
// case
}
```

当匹配到具体某个 case 时，i 即为 x 中该类型的值。  