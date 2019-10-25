package main

import "fmt"

//为类型添加方法

type Myint int

func (m Myint) IsZero() bool {
	return m == 0
}

func (m Myint) Add(other int) int {
	return other + int(m)
}

func main() {
	var b Myint

	fmt.Println(b.IsZero())

	b = 1

	fmt.Println(b.Add(2))
}
