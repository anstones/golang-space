package main

import (
	"fmt"
	"sync"
)

var (
	count int
	countGuard sync.Mutex
)

func GetCount() int {
	countGuard.Lock()

	// 函数退出时解除锁
	defer countGuard.Unlock()

	return count
}

func SetCount(c int)  {
	countGuard.Lock()
	count =c
	countGuard.Unlock()
}

func main()  {
	SetCount(100)
	fmt.Println(GetCount())
}