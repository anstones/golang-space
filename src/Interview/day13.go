package main

import "fmt"

func hello(i int) {
	fmt.Println(i)
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB() // 调用的是 People 自己的 ShowB 方法
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	i := 5
	defer hello(i) // hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5
	i = i + 10

	t := Teacher{}
	t.ShowA()
}
