package main

import (
	"fmt"
)

func Count(ch chan int, i int)  {
	ch <- i

}

func main() {
	chs := make([]chan int,10)
	for i :=0;i<10;i++{
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}

	for _,ch := range chs{
		if v, ok := <-ch; ok{
			fmt.Println(v)
		}
	}
}
