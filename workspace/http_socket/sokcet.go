package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "用法：%s IP 地址\n",os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil{
		fmt.Println("无效的IP地址")
	}else {
		fmt.Printf("IP地址是：%s\n",addr.String())
	}
	os.Exit(0)
}
