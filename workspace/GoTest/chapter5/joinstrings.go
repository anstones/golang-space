package main

// 可变参数数量不限定

import (
	"bytes"
	"fmt"
)

func joinStrings(slist ...string) string { // slist 为[]string 列表
	var b bytes.Buffer

	for _,s := range slist{
		b.WriteString(s)
	}

	return b.String()
}

func main()  {
	fmt.Println(joinStrings("pig ","and", " rat"))
	fmt.Println(joinStrings("hammer"," mom"," and"," hawk"))
}