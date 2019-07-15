package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Addr string
	Name string
}

var onlineMap map[string]Client

var message = make(chan string)

func main()  {
	listener,err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil{
		fmt.Println("net.Listen", err)
		return
	}

	defer listener.Close()

	go Manager()

	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("listerer.Accept:", err)
			continue
		}

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn)  {
	defer conn.Close()

	cliAddr := conn.RemoteAddr().String()

	cli := Client{make(chan string), cliAddr, cliAddr}

	onlineMap[cliAddr] = cli

	go WriteMsgToClient(cli, conn)

	message <- MakeStr(cli, "login")

	cli.C <- MakeStr(cli, "I am here")

	isQuit := make(chan bool)

	isTimeOut := make(chan bool)

	go func() {
		buf := make([]byte,1024)

		for {
			n,err:= conn.Read(buf)

			if n == 0{ // 对方断开
				isQuit <- true
				fmt.Println("conn.Read:", err)
				return
			}
			useMsg := string(buf[:n-1]) //windows nc 测试多个换行

			if len(useMsg) == 3 && useMsg == "who"{
				conn.Write([]byte("user list:\n"))
				for _,tmp := range onlineMap{
					useMsg = tmp.Addr +":" + tmp.Name +"\n"
					conn.Write([]byte(useMsg))
				}
			}else if len(useMsg) >= 8 && useMsg[:6] == "rename"{
				name := strings.Split(useMsg,"|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli

				cli.C <- MakeStr(cli,"rename successfully")
			}else {
				message <- MakeStr(cli, useMsg)

			}

			isTimeOut <- true
		}
	}()

	for {
		select {
		case <- isQuit:
			delete(onlineMap,cliAddr)
			message <- MakeStr(cli, "logout")
			return
		case <-isTimeOut:

		case <-time.After(10 * time.Second):
			delete(onlineMap,cliAddr)
			message <- MakeStr(cli,"timeout")
			return

		}
	}

}

func Manager()  {
	// 给map 分配空间
	onlineMap := make(map[string]Client)

	for {
		msg := <- message // 没有消息就组塞

		// 遍历map给每个成员发消息
		for _, cli := range onlineMap{
			cli.C <- msg  // 消息给信道
		}

	}
}

// 格式化字符串
func MakeStr(cli Client, str string) string  {
	return "[" + cli.Addr +"]" +cli.Name + ":" +str
 }

// 给当前用客户端发送消息
func WriteMsgToClient(cli Client, conn net.Conn)  {
	for msg := range cli.C{
		conn.Write([]byte(msg+"\n"))
	}
}