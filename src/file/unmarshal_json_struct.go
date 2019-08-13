package main

import (
	"encoding/json"
	"fmt"
)

// string --> json
// 解析到结构体，前提是知道被解析的josn结构体

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"Servers":[{"ServerName":"Local_Web", "ServerIP":"127.0.0.1"},{"ServerName":"Locel_DB","ServerIP":"172.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
