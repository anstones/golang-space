package main

import (
	"fmt"
)

// 数组和切片的区别：  数组需要明确指定大小，切片不需要。数组是值传递，切片是地址传递

func main() {

	array := [4]*string{new(string), new(string), new(string)}

	fmt.Printf("*array[0]: %s\n", *array[0])

	// 数组
	var x [3]int = [3]int{1, 2, 3}
	var y [3]int = x
	fmt.Println(x, y)
	y[0] = 999
	fmt.Println(x, y)

	//切片
	var a []int = []int{1, 2, 3}
	var b []int = a
	fmt.Println(a, b)
	b[0] = 999
	fmt.Println(a, b)
}
