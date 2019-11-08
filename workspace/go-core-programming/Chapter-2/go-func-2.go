package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

type Op func(int, int) int //主义一个函数类型，输入的是两个int 类型，返回位是一个in t 类型

func do(f Op, a, b int) int { //定义一个函数，第一个参数是函数类型Op
	return f(a, b) //函数类型交量可以直接用来进行函数调用
}

func main() {
	a := do(add, 10, 20)
	fmt.Println(a)
	s := do(sub, 10, 20)
	fmt.Println(s)
}
