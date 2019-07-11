package main

// go run -race  race  开启运行时对竞态问题的检测

import (
	"fmt"
	"sync/atomic"
)

// 竞态检测

var  (
	seq int64
)

func GenID() int64  {
	//atomic.AddInt64(&seq, 1)
	//	//return seq

	return atomic.AddInt64(&seq, 1)
}

func main()  {
	for i:= 0; i < 10; i++{
		go GenID()
	}

	fmt.Println(GenID())
}

