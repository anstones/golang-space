package main

import (
	"fmt"
)

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println(s, len(s), cap(s))
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))

	// var m map[string]int
	// m["a"]=1
	// if v := m["b"]; v != nil{
	// 	fmt.Println(v)
	// }
	m := make(map[string]int)
	m["a"] = 1
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}

	g := Work{3}
	var x A = g
	var y B = g
	// fmt.Println(y.ShowA()) //接口B 不包括方法 showA()
	// fmt.Println(x.ShowB()) //接口A 也不包括方法 ShowB()

	fmt.Println(x.ShowA())
	fmt.Println(y.ShowB())
}
