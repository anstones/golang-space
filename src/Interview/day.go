package main

import "fmt"

type Person struct {
	name string
}

func hello(num ...int) {
	num[0] = 18
}

func main() {
	var m map[Person]int
	p := Person{"mike"}
	fmt.Println(m[p])
	// 打印一个map中不存在的值时,返回元素类型的零值。
	// 这个例子中，m 的类型是 map[person]int，因为 m 中不存在 p，所以打印 int 类型的零值，即 0

	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}

func main2() {
	var ch chan int
	ch := make(chan int)
	<-ch // 读取channel
	//ch <-  // 写 channel 是必须带上值,
}
