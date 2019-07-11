package main

import (
	"fmt"
	"time"
)

func main()  {
	exit := make(chan int)

	fmt.Println("start")

	// 一秒后调用函数
	time.AfterFunc(time.Second, func() {
		fmt.Println("one second after")

		exit <- 0
	})

	<-exit
}
