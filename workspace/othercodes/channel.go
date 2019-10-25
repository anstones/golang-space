package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    numberGoroutines = 4   //使用goroutine的数量
    taskLoad         = 10    //要处理的工作数量
)

var wg sync.WaitGroup

func init() {
    rand.Seed(time.Now().Unix())
}

func main() {
    tasks := make(chan string, taskLoad)
    wg.Add(numberGoroutines)

   //启动goroutine来处理工作
    for gr := 1; gr <= numberGoroutines; gr++ {
        go worker(tasks, gr)
    }

    //增加一组要处理的工作
    for post := 1; post <= taskLoad; post++ {
        tasks <- fmt.Sprintf("Task:%d", post)
    }

    //指派完要处理的工作后关闭通道
    //所有的goroutine根据通道关闭的条件退出
    close(tasks)

   //等待所有工作完成
    wg.Wait()
}


func worker(tasks chan string, worker int) {
    //通知main函数goroutine返回
    defer wg.Done()

    for {
       //等待要处理的工作，即分配工作
       //通道中没有工作时，goroutine返回
        task, ok := <-tasks
        if !ok {
            //通道为空，已经关闭
            fmt.Printf("worker:%d :shutting Down\n", worker)
            return
        }

        //goroutine开始处理分配到的一个工作
        fmt.Printf("worker:%d:started %s\n", worker, task)

        //随机等待一段时间模拟处理工作
        sleep := rand.Int63n(100)
        time.Sleep(time.Duration(sleep) * time.Millisecond)

        //显示完成了某个工作
        fmt.Printf("worker:%d:completed %s\n", worker, task)
    }
}
