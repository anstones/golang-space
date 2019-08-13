package main

import (
	"encoding/json"
	"fmt"
)

// json --> string

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "local_Web", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "local_DB", ServerIP: "172.0.0.2"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
