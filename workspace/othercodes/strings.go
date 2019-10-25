package main

import (
	"strings"
)

func main() {
	str := "你好世界，这个世界真好。"
	new := "地球"
	old := "世界"
	n := -1 // 替换字符串中所有符合的字符串
	println(strings.Replace(str, old, new, n))
}
