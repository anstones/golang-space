package main

// channel 超时

import (
	"fmt"
	"time"
)

var timeout = make(chan bool, 1)
var ch = make(chan int,10)

func main() {
	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()

	select {
	case <-ch:
		if v,ok := <-ch;ok{
			fmt.Println(v)
		}
	case <-timeout:
		fmt.Println("timeout")
	}
}
