package telentecho

import (
	"fmt"
	"strings"
)

func processTelnetCommand(str string, exitchan chan int) bool {
	 // telent   键盘输入@close 表示终止本次会话
	 if strings.HasPrefix(str, "@cclose"){
	 	fmt.Println("session closed")

	 	return false
	 }else if strings.HasPrefix(str, "shutdown"){ //telent   键盘输入@shutdown 表示终止服务器进程
		 fmt.Println("server shutdown")

		 // 往通道写入0 阻塞等待接收方处理
		 exitchan <- 0

		 return false
	 }

	 fmt.Println(str)

	 return true
}