package main

import "fmt"

//defer 语句必须先注册后才能执行，如果defer 位于return 之后，则由fer 因为没有注册，不会执行。

func main() {
	defer func() {
		fmt.Println("first")
	}()
	a := 0
	fmt.Println(a)
	return
	defer func() {
		fmt.Println("second")
	}()
}
