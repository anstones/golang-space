package main

import (
	"fmt"
	"sort"
)

type stuScore struct {
	name  string
	score int
}

type stuScores []stuScore

func (s stuScores) Len() int {
	return len(s)
}

func (s stuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//func (s stuScores) Less(i,j int) bool {
//	return s[i].score > s[j].score   //逆序排序
//}

func (s stuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := stuScores{
		{"张三", 95},
		{"李四", 91},
		{"王五", 96},
		{"赵六", 90},
	}

	fmt.Println("=======排序前=======")
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}
	// 排序
	sort.Sort(stus)

	fmt.Println("=======排序后=======")
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}

	//判断是否排序
	fmt.Println("是否排序？ \n", sort.IsSorted(stus))
}
