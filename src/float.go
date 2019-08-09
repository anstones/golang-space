package main

import "fmt"

func main() {
	var value1 float64 // float32 :小数点后7位 ； float64:小数点后15位

	value1 = 1
	value2 := 1.000000000000001

	if value1 == value2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}
}
