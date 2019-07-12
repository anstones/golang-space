package main

// 因为循环没有结束条件， 程序会不停的创建goroutine

import (
	"fmt"
	"runtime"
)

func consumer(ch chan int)  {
	for {
		data := <- ch

		fmt.Println(data)
	}
}

func main() {
	ch := make(chan int)

	for {
		var dummy string

		fmt.Scan(&dummy)

		go consumer(ch)

		fmt.Println("goroutine: ", runtime.NumGoroutine())
	}
}
