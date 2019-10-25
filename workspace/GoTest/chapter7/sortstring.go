package main

import (
	"fmt"
	"sort"
)

// 将[]string 定义为MyStringList 类型
type MyStringList []string

// 实现sort.interface 接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// 实现sort.interface 接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

//实现sort.interface 接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	sort.Sort(names)

	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
