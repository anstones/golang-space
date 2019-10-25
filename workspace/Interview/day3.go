package main

import "fmt"

func main() {
	li := make([]int, 5)
	li = append(li, 1, 2, 3)
	fmt.Println(li)

	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	fmt.Println(SumHui(1, 5))
}

func SumHui(x, y int) (sum int, err error) {
	return x + y, nil
}

// new() 与 make() 的区别?
//new用于结构对象生成 make用于map 切片生成
