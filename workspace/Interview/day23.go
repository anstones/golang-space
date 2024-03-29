package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4 //当使用 s1[1:] 获得切片 s2，和 s1 共享同一个底层数组，这会导致 s2[1] = 4 语句影响 s1
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7) //而 append 操作会导致底层数组扩容，生成新的数组，因此追加数据后的 s2 不会影响 s1。
	fmt.Println(s1)

	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
}
