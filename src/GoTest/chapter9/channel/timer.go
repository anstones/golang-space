package main

import (
	"fmt"
	"time"
)

func main()  {
	// 创建一个打点器，每500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond*500)

	// 计时器，2秒触发一次
	stoper := time.NewTimer(time.Second*2)

	var i int

	for {
		select {
		case <-stoper.C:
			fmt.Println("stop")
			goto StopHere
		case <-ticker.C:
			i ++
			fmt.Println("tick ", i)

		}
	}

	StopHere:
		fmt.Println("done")
}
