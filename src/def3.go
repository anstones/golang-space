package main

import (
	"fmt"
)

type add func(a int, b int) int

func simple() func(a,b int ) int  {
	f := func(a,b int) int{  //将匿名函数传递到simple
		return a+b
	}
	return f

}

func main() {
	//var a add = func(a int, b int) int {
	//	return a + b
	//}
	//fmt.Printf("%T\n",a)
	//s := a(5, 6)
	//fmt.Println("Sum", s)

	s := simple()
	fmt.Printf("%T\n", s)
	fmt.Println(s(60,7))


}