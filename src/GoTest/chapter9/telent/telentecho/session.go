package telentecho

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)


// 链接的会话逻辑
func handleSession(conn net.Conn, exitchan chan int)  {
	fmt.Println("session start")

	// 创建网络连接 读取器
	reader := bufio.NewReader(conn)

	for {
		str,err := reader.ReadString('\n')
		if err == nil{
			// 去掉字符串尾部的空格
			str = strings.TrimSpace(str)

			// 处理Telent
			if !processTelnetCommand(str, exitchan){
				conn.Close()
				break
			}

			// echo 逻辑，发什么数据原样返回
			conn.Write([]byte(str + "\r\n"))
		}else {
			fmt.Println("session closed")
			conn.Close()
			break
		}
	}
}