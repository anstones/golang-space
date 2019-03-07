package main

import (
	"fmt"
	"unsafe"
)


var a = 10
var name = "xiaom"
var A,B,C = 1,"www.oeasy.cn",false
var x,y = 23,56
var m,n = 5,7

const (
    aa = "abc"
    bb = len(aa)
    c = unsafe.Sizeof(aa)
)

func main() {
	fmt.Println("hello,go")
	fmt.Println("yihaizuc")
	var b bool = false
	top := 789
	x,y = y,x
	var _, f = 5,n
	if b{
		fmt.Println(a, name, A, B, C, top, x,y,f)
	}else{
		_,nums,strs := numbers()
		fmt.Println(nums,strs)
		fmt.Printf("面积为 %d\n", c)
	}
	
}


func numbers()(int,int,string) {
	p,q,v := 1111,222,"12356"
	return p,q,v
}
