package main

import (
	"fmt"
)

type Personer interface {
	SayHello()
}
type Student struct {
	Name string
}

func (s *Student) SayHello() {
	fmt.Println("Hello,", s.Name)
}
func main() {
	var p Personer = Student{Name: "zhangsan"}
	p.SayHello()
}
