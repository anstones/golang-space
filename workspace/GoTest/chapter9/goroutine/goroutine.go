package main

// 如果需要在goroutine中返回数据  用channel把数据返回

import (
	"fmt"
	"time"
)

func Runing()  {
	var times int

	for {
		times ++
		fmt.Println("tick", times)

		time.Sleep(1*time.Second)
	}
}

func main()  {
	go Runing()

	var input string
	fmt.Scanln(&input)
}

