package main

import (
	"GoTest/chapter9/telent/telentecho"
	"os"
)

func main()  {
	exchan := make(chan int )

	// 并发运行服务器
	go telentecho.Server("127.0.0.1:7001", exchan)

	// 通道阻塞，等待返回值
	code := <-exchan

	// 标记程序返回值并推出
	os.Exit(code)
}