package main

import (
	"fmt"
	"sync"
)

var x int //一个全局的变量

func add(wg *sync.WaitGroup, m *sync.Mutex){
	m.Lock()  // 用Mutex处理竞态条件  类似于python线程锁
	x += 1
	m.Unlock()
	wg.Done()
}

func increment(wg *sync.WaitGroup, ch chan bool)  {
	ch <- true // 用缓冲信道处理竞态条件
	x+=1
	<-ch  // 缓冲信道被消费
	wg.Done()
}

func main()  {
	var wg sync.WaitGroup
	//var m sync.Mutex
	var ch = make(chan bool,1)
	//for i :=0;i<1000;i++{
	//	wg.Add(1)
	//	go add(&wg, &m)
	//}
	for i :=0;i<1000;i++{
		wg.Add(1)
		go increment(&wg,ch)
	}
	wg.Wait()
	fmt.Println("final value of x", x)
}