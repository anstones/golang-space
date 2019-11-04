package main

import "fmt"

const (
	a = 1 << iota
	b = 1 << iota
	c = 3
	d = 1 << iota //左移 n 位就是乘以 2 的 n 次方
)

func main() {
	fmt.Println(a, b, c, d)
}
