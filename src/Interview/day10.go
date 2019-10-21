package main

import "fmt"

func main() {
	a := [2]int{4, 5}
	b := [3]int{4, 5}
	if a == b { //Go 中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的组成部分，所以a和b是不同的类型，是不能比较的，所以编译错误。
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	x := [5]int{1, 2, 3, 4, 5}
	y := x[3:4:4]
	//第三个参数 k 用来限制新切片的容量，但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i。
	// y 的 长度(len=j-i) 4-3 = 1， 容量(cap=k-i) 4-3 =1， y的值 [4]
	fmt.Println(y[0])

	m := 5
	n := 8.1
	fmt.Println(m + n) // 类型不同

}
