package main

// 循环从信道接受数据

import (
	"fmt"
	"time"
)

func main()  {

	ch := make(chan  int)


	go func() {
		for i:=1; i<=3;i++{
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("wait goroutine")

	for data := range ch{
		fmt.Println(data)

		if data == 3{  // 当data为3时  goroutine已经结束，即没有发送通道，此时必须要终止循环
			break
		}
	}
	fmt.Println("all done")
}