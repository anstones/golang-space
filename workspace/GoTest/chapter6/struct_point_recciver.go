package main

import "fmt"

// 指针类型接收器： 调用方法，修改接收器指针的任意成员变量，在方法结束后，修改有效
// 更接近与 self和this的概念

type Property struct {
	value int
}

func (p *Property) Setvalue(v int) {
	p.value = v
}

func (p *Property) Getvalue() int {
	return p.value
}

func main() {
	// 实例化属性
	p := new(Property)
	//设置值
	p.Setvalue(10)

	// 取值
	fmt.Println(p.Getvalue())
}
