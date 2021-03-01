## 1，请指出以下函数的调用顺序  

```go
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int { return c }
func g() int  { return a }
func sqr(x int) int { return x*x }
func u() int { return 1}
func v() int { return 2}
```

### 答案  

我在这个题之前，总结了一个知识点：Go语言规范规定了求值顺序，其中有一条：在包级别，初始化依赖决定了变量声明中表达式的求值顺序。  

这里的关键是依赖。我们一步步分析。  

var a, b, c = f() + v(), g(), sqr(u()) + v() 这里看起来应该先给 a 赋值，也就是要调用 f() 和 v() 这两个函数，但是 f 函数依赖了 c，所以需要先初始化 c，也就是调用 sqr(u())+v()，这一个表达式又应该先调用 u()，然后是 sqr()，接着是 v()，这样 c 初始化完了，所以，f() 可以调用了（v 会再调用一次），最后才是 g() 来初始化 b。  

所以最后的顺序是：u()、sqr()、v()、f()、v()、g()。  

其实这和包依赖初始化顺序是类似的，可以看我之前发的包初始化的图。  



## 2，有如下代码：  

```go
type MyWriter struct{}

func (m *MyWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

var _ io.Writer = (*MyWriter)(nil)  //(*MyWriter)(nil)  这是类型转换
```

请问，var _ io.Writer = (*MyWriter)(nil) 有什么用？