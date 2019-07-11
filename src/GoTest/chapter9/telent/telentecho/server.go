package telentecho

import (
	"fmt"
	"net"
)

func Server(addr string, exitchan chan int)  {

	// 监听
	l, err := net.Listen("tcp", addr)

	// 如果监听发生错误，则推出
	if err != nil{
		fmt.Println(err)
		exitchan <- 1
	}

	fmt.Println("listen : ", addr)

	// 延迟关闭监听器
	defer l.Close()

	// 循环
	for {
		// 新链接没到来之前，Accept是阻塞的
		conn, err := l.Accept()

		if err != nil{
			fmt.Println(err)
			continue
		}

		go handleSession(conn, exitchan)
	}
}
