package main

//无缓冲通道

import "fmt"

func Printer(c chan int)  {
	for {
		data :=  <- c

		if data == 0{
			break
		}

		fmt.Println(data)
	}

	//通知main已经结束循环
	c <- 0
}

func main()  {
	c := make(chan int)

	go Printer(c)

	for i := 1;i <= 10;i++{
		c<- i
	}

	// 通知Printer  结束循环发送数据
	c<-0

	// 等待printer 结束
	<-c
}