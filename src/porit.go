package main

import (
	"fmt"
)

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)  // 指针的引用


	c := 255
	d := &c
	fmt.Println("address of b is", d)
	fmt.Println("value of b is", *d)  // 指针的解引用

	*d++  // 利用指针修改值
	fmt.Println("new value of b is", c)
}