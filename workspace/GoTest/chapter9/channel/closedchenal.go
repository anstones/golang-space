package main

//从已关闭的通道取值 不会阻塞

import "fmt"

func main()  {

	ch := make(chan int, 3)


	ch <- 0

	ch <- 1

	close(ch)


	for i := 0; i< cap(ch); i++{
		v,ok := <-ch

		fmt.Println(v,ok)
	}



}
