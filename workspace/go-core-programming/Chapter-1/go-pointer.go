package main

import "fmt"

var a = 11

type User struct {
	Name string
	Age  int
}

func main() {
	as := &a         //*as 和a 的位都是11
	fmt.Println(as)  //指针地址
	fmt.Println(*as) // 指针的值

	andes := User{
		Name: "andes",
		Age:  20,
	}
	p := &andes
	fmt.Println(p.Name) //. name 通过"."操作符访问成员交量
}
