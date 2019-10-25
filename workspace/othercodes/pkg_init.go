package main

import (
	"fmt"
	"runtime"
)


func init() {
	fmt.Printf("%v\n", m)
	info = fmt.Sprintf("OS:%s, Arch:%s", runtime.GOOS, runtime.GOARCH)
}

func outerfunc() {
	defer fmt.Println("sfsfs")
	fmt.Println("asfaaaaaaaaaaaa")
	
}

var m = map[int] string{1:"A",2:"B",3:"C"}
var info string

func main() {
	fmt.Println(info)
	outerfunc()
}


