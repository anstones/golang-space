package main

import (
	"fmt"
	"net"
	"time"
)

func socketRecv(conn net.Conn, exitchan chan string)  {
	buff := make([]byte, 1024)

	for {
		_, err := conn.Read(buff)

		if err != nil{
			break
		}
	}

	exitchan <- "recv exit"  // 函数结束，向通道发送通知
}

func main() {
	conn,err := net.Dial("tcp", "www.163.com:80")

	if err != nil{
		fmt.Println(err)
		return
	}

	exit := make(chan string)

	go socketRecv(conn, exit)

	time.Sleep(time.Second)  //接受需要一个过程，所以等待1秒

	conn.Close()  //主动关闭套接字，此时会触发套接字链接错误

	fmt.Println(<-exit) //从通道接收退出数据，也就是等待接收goroutine结束


}