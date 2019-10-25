package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkErr(err error)  {
	if err != nil{
		fmt.Printf("错误：%s",err.Error())
		os.Exit(1)
	}
}

func main() {
	sevvice := os.Args[1]
	tcpaddr, err := net.ResolveTCPAddr("tcp4", sevvice)
	checkErr(err)
	conn,err := net.DialTCP("tcp", nil, tcpaddr)
	checkErr(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	result,err := ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Println(string(result))
	os.Exit(0)
}
