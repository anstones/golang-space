package main

// 指针传递使得多个函数可以操作同一个对象
// 指针传递轻量级（传的只是内存位置）
// 节省内存（没有渎职变量的值），而且赋予了函数修改外部变量的能力， 所以不必return

import "fmt"

func main() {
	x := 3
	fmt.Println("x = ", x, "&x = ", &x)

	y := add(x)                       // 执行add函数，但实际上修改的是x副本
	fmt.Println("x = ", x, "y = ", y) // 输出x还是原来的值， 修改的是副本y的值

	z := addp(&x)                     // 实际上是修改了x的值
	fmt.Println("x = ", x, "z = ", z) // 输出的X已经被修改
	fmt.Println("&x = ", &x)          // x的地址还是原来的地址
}

func add(x int) int {
	x++
	return x
}

func addp(a *int) int {
	*a++
	return *a
}
