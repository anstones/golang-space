package main

import "fmt"

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

type SB struct {
	m string
}

func n() *SB {
	return &SB{"foo"}
}

func main() {
	var x []int  //声明的是 nil 切片, 不会分配内存，优先选择
	a := []int{} //声明的是长度和容量都为 0 的空切片
	fmt.Println(x, a)

	s := S{}
	p := &s
	f(s) //A
	// g(s) //B 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。
	f(p) //C
	// g(p) //D 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。

	y := n()
	fmt.Println(y.m)
}
