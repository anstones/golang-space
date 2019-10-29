package main

import (
	"fmt"
)

var a bool = true

func main() {
	// var x string = nil //cannot use nil as type string in assignment
	var x string = ""
	if x == "" {
		x = "default"
	}
	fmt.Println(x)

	defer func() {
		fmt.Println("1")
	}()
	if a == true {
		fmt.Println("2")
		return
	}
	// return 之后的 defer 是不能注册的， 也就不能执行后面的函数或方法
	defer func() {
		fmt.Println("3")
	}()

}
