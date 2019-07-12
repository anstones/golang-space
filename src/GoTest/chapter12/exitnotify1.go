package main

// 不使用通道，使用等待组  优化同步

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func socketRecv(conn net.Conn, wg *sync.WaitGroup)  {
	buff := make([]byte, 1024)

	for {
		_, err := conn.Read(buff)

		if err != nil{
			break
		}
	}

	wg.Done()  // 函数结束，发送通知
}

func main() {
	conn,err := net.Dial("tcp", "www.163.com:80")

	if err != nil{
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go socketRecv(conn, &wg)

	time.Sleep(time.Second)  //接受需要一个过程，所以等待1秒

	conn.Close()  //主动关闭套接字，此时会触发套接字链接错误

	//等待goroutine退出完毕
	wg.Wait()
	fmt.Println("recv done")


}
