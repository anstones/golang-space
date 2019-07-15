package main


// 优雅的解决tcp黏包

import (
	"GoTest/chapter12/tcppkt/tcppage"
	"net"
	"strconv"
)

func main() {
	// 测试次数
	const TestCount  =100000

	const address  = "127.0.0.1:8010"

	// 接收到的计数器
	var recvCounter int

	acceptor := tcppage.NewAcceptor()

	acceptor.Start(address)

	acceptor.OnSessionData = func(conn net.Conn, data []byte) bool{

		str := string(data)

		// 转为数字
		n,err := strconv.Atoi(str)

		// 接收错位，其他错误
		if err != nil || recvCounter != n {
			panic("failed")
		}

		recvCounter ++

		if recvCounter >= TestCount{
			acceptor.Stop()
			return false
		}

		return true
	}

	tcppage.Connector(address, TestCount)

	acceptor.Wait()

}