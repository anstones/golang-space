package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func add(args ...int) int {

	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	str := "hello"
	//str[0] = 'x'  // 常量，Go 语言中的字符串是只读的,此处报错
	fmt.Println(str)

	p := 1
	incr(&p)
	fmt.Println(p)

	fmt.Println(add(1, 3, 7))
}
