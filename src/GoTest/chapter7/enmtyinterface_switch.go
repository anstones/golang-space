package main

import "fmt"

// 批量判断空接口中数据的类型 用switch

func printTpye(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}

func main() {
	printTpye(1024)
	printTpye("pig")
	printTpye(true)
}
