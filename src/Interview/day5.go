package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 { //map属于不可比较类型
		fmt.Println("sm1 == sm2")
	}
	//如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，
	// 比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
	// 可比较的有：  bool、数值型、字符、指针、数组等，
	// 不可比较的有：切片、map、函数等。
}
