package main

// 闭包

import "fmt"

func Accumulate(value int)func() int {
	// 返回一个闭包
	return func() int {
		value ++
		// 返回累加值
		return value
	}
}

func main()  {
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())


	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
}