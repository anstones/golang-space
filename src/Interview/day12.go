package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People // 嵌套结构体 people结构体为内部结构体
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB") // 定义和内部结构体相同的方法，内部类型的方法就会被“屏蔽”
}

func main() {
	i := -5
	j := +5
	fmt.Printf("%+d %+d \n", i, j)

	t := Teacher{}
	t.ShowB()
}
