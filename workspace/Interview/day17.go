package main

import (
	"fmt"
)

func Yes() (str string, err error) {
	return "yes", nil
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type work struct {
	i int
}

func (w work) ShowA() int {
	return w.i + 10
}

func (w work) ShowB() int {
	return w.i + 20
}

func main() {
	var x string
	// x, _ := Yes()  //x 已经声明，不能使用 :=
	x, _ = Yes()
	fmt.Println(x)
	x, y := Yes() // 当多值赋值时，:= 左边的变量无论声明与否都可以
	// x, y = Yes()  // y 没有声明
	fmt.Println(x, y)

	fmt.Println(increaseA())
	fmt.Println(increaseB())

	var a A = work{3}
	s := a.(work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
