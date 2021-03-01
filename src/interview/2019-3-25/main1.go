package main

func main() {
	Sum(1, 2)
	len("Hello") //编译不过
}

func Sum(a, b int) int {
	return a + b
}
