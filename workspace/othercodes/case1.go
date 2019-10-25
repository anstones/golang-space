package main

import (
	"fmt"
)

var x interface{}

func main() {

	x = 1
	switch i := x.(type) { // 类型判断switch
	case nil:
		fmt.Printf("x的类型为%T", i)
	case int:
		fmt.Printf("x的类型为%T", i)
	case float64:
		fmt.Printf("x的类型为%T", i)
	case bool:
		fmt.Printf("x的类型为%T", i)
	case string:
		fmt.Printf("x的类型为%T", i)
	default:
		fmt.Printf("x为未知类型")
	}

}
