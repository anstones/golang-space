package main

import "fmt"

func main() {
	str := "123" + "abc"
	fmt.Println(str)
	fmt.Sprintf("abc%d", 123)

	const (
		x = iota
		_
		y
		z = "zz"
		k
		p = iota
	)
	fmt.Println(x, y, z, k, p)

	//var x = nil  // 报错
	//var x interface{} = nil
	//var x string = nil  //报错
	//var x error = nil

	// nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。
	// 强调下 D 选项的 error 类型，它是一种内置接口类型，

}
