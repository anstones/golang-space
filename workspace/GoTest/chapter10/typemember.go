package main

import (
	"fmt"
	"reflect"
)

func main() {
	type cat struct {
		Name string
		Type int     `json:"type" id:"100"`
	}

	// 结构体实例化
	ins := cat{"mimi", 12}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	// typeOfCat.NumField() 返回结构体字段数量
	for i:= 0; i< typeOfCat.NumField(); i++{

		// 获取每个成员的结构体字段类型
		fileType := typeOfCat.Field(i)

		fmt.Printf("name:%v tag: %v\n", fileType.Name, fileType.Tag )

	}

	// 通过字段名，找到字段类型的信息
	if catType, ok := typeOfCat.FieldByName("Type");ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
