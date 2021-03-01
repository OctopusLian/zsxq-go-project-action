Go 语言规范中定义了求值顺序。（描述改进）  

1、一般地，在计算表达式、赋值或返回语句的操作数时，所有函数调用、方法调用、通讯操作（\<-）等在词法层面按照从左到右的顺序求值。注意以下代码中的注释。  

```go
a := 1
f := func() int { a++; return a }
x := []int{a, f()}            // x may be [1, 2] or [2, 2]: evaluation order between a and f() is not specified
m := map[int]int{a: 1, a: 2}  // m may be {2: 1} or {2: 2}: evaluation order between the two map assignments is not specified
n := map[int]int{a: f()}      // n may be {2: 3} or {3: 3}: evaluation order between the key and the value is not specified
```

以上代码的总结是：求值表达式中  

1）变量和函数的先后顺序未定义；  
2）map 元素的先后顺序未定义；  
3）map 中 key 和 value 的先后顺序未定义；  

2、然而，在包级别，初始化依赖决定了变量声明中表达式的求值顺序。  

只要是顺序未定义的，在实际项目中一定要避免使用，否则可能出现莫名其妙的“bug”，被“坑”~  

