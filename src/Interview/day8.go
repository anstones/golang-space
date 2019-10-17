package main

import "fmt"

func hello() []string {
	return nil
}

func GetValue() int {
	return 1
}

func main() {
	h := hello // 是将 hello()函数赋值给变量 h，而不是函数的返回值，所以输出 not nil
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	//i := GetValue()
	var i interface{}
	switch i.(type) { // 类型选择的语法形如：i.(type)，其中i是接口，type是固定关键字；需要注意的是，只有接口类型才可以使用类型选择。
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
