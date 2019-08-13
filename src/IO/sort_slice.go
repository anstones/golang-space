package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s) // 升序
	fmt.Println(s)

	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s) //降序

	sort.Ints(s)
	fmt.Println(sort.SearchInts(s, 3)) //查询3是否在s中，如果存在，返回3所在下表
}
