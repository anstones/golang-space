package main

import (
	"fmt"
	"runtime"
)

func say(s string)  {
	for i:= 0;i<5;i++{
		runtime.Gosched()
		//time.Sleep(time.Second)
		fmt.Println(s)
	}
}

func main()  {
	fmt.Println("start go")
	go say("world")
	say("hello")
}
