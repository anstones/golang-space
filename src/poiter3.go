package main

import (
	"fmt"
)

func Change(val *int) {
	*val = 55
}
func main() {
	a := 58
	fmt.Println("value of a before function call is",a)
	//b := &a
	//Change(b)
	Change(&a)  // 传递指针参数
	fmt.Println("value of a after function call is", a)
}