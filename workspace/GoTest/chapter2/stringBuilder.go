package chapter2

import (
	"bytes"
	"fmt"
	"math"
	"unicode/utf8"
)

func main()  {

	tip1 := "genji is ninja"
	tip2 := "忍者"
	fmt.Println(len(tip1))  // ASCII 字符串计算长度
	fmt.Println(utf8.RuneCountInString(tip2)) // Unicode 计算长度


	theme := "忍者阻击start"
	for _,s := range theme{
		fmt.Printf("unicode: %c\n", s)  //ASCII 遍历用下标，unicode 遍历用for range
	}


	hammer := "吃我一锤"
	sickle := "死吧"
	var stringBuilder bytes.Buffer
	// 高效的字符串拼接
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	fmt.Println(stringBuilder.String())


	var progress = 2
	var target = 8
	title := fmt.Sprintf("已采集%d个药草，还需要%d个完成任务",progress,target)
	fmt.Println(title)
	pi := math.Pi
	variant := fmt.Sprintf("%v %v %v", "月球基地",pi,true)
	fmt.Println(variant)
}
