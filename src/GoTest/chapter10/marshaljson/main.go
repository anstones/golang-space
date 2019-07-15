package main

// 模拟json.marshal

import (
	"GoTest/chapter10/marshaljson/json"
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
			{Name:"Roll and roll",Level:3},
			{Name:"Flash your dog eye",Level: 6},
			{Name:"Time to have Lunch",Level:3},
		},
	}

	if result,err := json.MarshalJson(a); err == nil{
		fmt.Println(result)

	}else {
		fmt.Println(err)
	}
}