package main

import "fmt"

var (
	//size := 1024 // 编译不通过
	size     = 1024
	max_size = size * 2
)

func main() {
	//s := new([]int)  // 编译不通过，需要用make
	s := make([]int, 0)
	s = append(s, 1)
	fmt.Println(s)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//s1 = append(s1,s2)
	s1 = append(s1, s2...)
	fmt.Println(s1)

	fmt.Println(size, max_size)
}
