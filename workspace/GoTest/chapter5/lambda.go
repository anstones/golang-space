package main

import (
	"flag"
	"fmt"
)

var skillParm = flag.String("skill", "", "skill to peoform")
// flag 包： 解析命令行参数
//第一个参数 ：flag名称为flagname
//第二个参数 ：flagname默认值为1234
//第三个参数 ：flagname的提示信息

func main()  {
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f,ok := skill[*skillParm]; ok{
		f()
	}else{
		fmt.Println("skill mot fond")
	}
}
