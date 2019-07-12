package main

// 模拟json.marshal

import (
	"fmt"
)

func main() {
	// 声明技能结构
	type Skill struct {
		Name string
		Level int
	}

	//声明角色结构
	type Actor struct {
		Name string
		Age int
		Skills []Skill
	}

	//填充基本角色数据
	a := Actor{
		Name:"cow boy",
		Age:30,
		Skills:[]Skill{
			{"Roll and roll",3},
			{"Flash your dog eye", 6},
			{"Time to have Lunch",3},
		},
	}

	if result,err := MarshalJson(a); err == nil{
		fmt.Println(result)

	}else {
		fmt.Println(err)
	}
}