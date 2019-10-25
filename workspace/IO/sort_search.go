package main

//search :搜索元素x是否存在于已经升序排序号的切片中

import (
	"fmt"
	"sort"
)

func GuessingGame() {
	var s string
	fmt.Printf("选择一个从0到100的整数：\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("这个数字是否小于等于%d？", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("你选择的数字是%d \n", answer)
}

func main() {
	x := 11
	s := []int{3, 6, 8, 11, 45}
	pos := sort.Search(len(s), func(i int) bool {
		return s[i] >= x
	})
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, "在 s中的位置是：", pos)
	} else {
		fmt.Println("s不包含元素 ", x)
	}

	GuessingGame()
}
