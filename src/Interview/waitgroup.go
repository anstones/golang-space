package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Main start", time.Now())

	go func() {
		fmt.Println("go start", time.Now())
		fmt.Println("go end", time.Now())
	}()

	fmt.Println("Main end", time.Now())

	main2()
	//waitgroup()
}

func main2() {
	fmt.Println("Main start", time.Now())

	go func() {
		fmt.Println("go start", time.Now())
		fmt.Println("go end", time.Now())
	}()
	time.Sleep(1 * time.Second) //通过超时等待goroution 结束
	fmt.Println("Main end", time.Now())
}

func waitgroup() {
	wg := sync.WaitGroup{}
	fmt.Println("Main start", time.Now())
	wg.Add(1) // 计数器增长1
	go func() {
		fmt.Println("go start", time.Now())
		wg.Done() // 计数器减少1
		fmt.Println("go end", time.Now())
	}()
	wg.Wait()
	fmt.Println("Main end", time.Now())
}
