package main

// 非指针类型接收器   Add方法类似于只读的方法，Add方法不会对成员进行任何修改

import "fmt"

type Ponit struct {
	x int
	y int
}

func (p Ponit) Add(a, b int) Ponit {
	return Ponit{p.x + a, p.y + b}
}

func main() {
	p1 := Ponit{1, 2}

	a := 3
	b := 4

	result := p1.Add(a, b)

	fmt.Println(result)

	fmt.Printf("%v\n", p1)

}
