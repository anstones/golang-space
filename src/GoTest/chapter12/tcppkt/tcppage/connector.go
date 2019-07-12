package tcppage

import (
	"fmt"
	"net"
	"strconv"
)

// 链接器


// 连接器，传入链接地址和发送封包次数
func Connector(addr string, sendTimes int)  {

	conn, err := net.Dial("tcp", addr)

	if err != nil{
		fmt.Println(err)
		return
	}

	// 循环指定次数
	for i:=0; i< sendTimes; i++{
		str:= strconv.Itoa(i)

		// 发送封包
		if err := WritePacket(conn, []byte(str)); err != nil{
			fmt.Println(err)
			return
		}
	}
}

