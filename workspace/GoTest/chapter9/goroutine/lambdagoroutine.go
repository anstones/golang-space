package main

// 所有的goroutine在main函数结束后会一同结束
// 暂时没有接口获取goroutine的ID
// Go1.5 以后默认执行 runtime.GOMAXPROCS(runtime.NumCPU())

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	go func() {
		var times int
		for {
			times ++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}

	}()

	runtime.GOMAXPROCS(runtime.NumCPU())  //指定运行的CPU数量

	var input string
	fmt.Scanln(&input)
}
