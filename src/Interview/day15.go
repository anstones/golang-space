package main

import "fmt"

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
	var s1 []int
	fmt.Println(s1)
	var s2 = []int{}
	fmt.Println(s2)
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}

	i := 65
	fmt.Println(string(i))

}
