package main

import "fmt"

// 缓冲通道
// 缓冲通道被填满，继续发送数据到通道会发生阻塞
// 缓冲通道为空，尝试从中取数据 发生阻塞

func main()  {

	ch := make(chan int,3)

	fmt.Println(len(ch))

	ch <- 1
	ch <- 2
	ch <- 3


	fmt.Println(len(ch))
}
