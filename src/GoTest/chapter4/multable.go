package main

import "fmt"

func main()  {

	for x := 1;x <= 9; x++{
		for y :=1; y<=x;y ++{
			fmt.Printf("%d*%d = %d ", y,x,x*y)
		}
		fmt.Println()
	}

	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v:= range c{
		fmt.Println(v)  // for range 遍历channel 只返回通道内的值
	}

	var s = "hello"
	switch  {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")

	}
}
