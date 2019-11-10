package main

import "fmt"

// defer 函数的实参在注册时通过值拷贝传递进去。下面示例代码中， 实参a 的值在defer 注
// 册时通过值拷贝传递进去，后续语句a＋＋并不会影响defer i吾句最后的输出结果。

func f() int {
	a := 0
	defer func(i int) {
		fmt.Println("defer i =", i)
	}(a)
	a++
	return a
}

func main() {
	f()
}
