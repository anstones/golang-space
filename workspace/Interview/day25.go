package main

import "fmt"

type MyInt int

func (i MyInt) PrintInt() {
	fmt.Println(i)
}

// func (i int) PrintInt() {
// 	fmt.Println(i)
// }

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	// var i int = 1
	var i MyInt = 1
	i.PrintInt()

	// 基于类型创建的方法必须定义在同一个包内，上面的代码基于 int 类型创建了 PrintInt() 方法，
	// 由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错。

	// var peo People = Student{}
	var stu = Student{}
	think := "speak"
	fmt.Println(stu.Speak(think))
}
