package main

//使用反射访问结构体的成员字段的值

import (
	"fmt"
	"reflect"
)

type dummy struct {
	a int
	b string
	float32
	bool

	next *dummy
}


func main() {
	// 值包装结构体
	d := reflect.ValueOf(dummy{next: &dummy{},})

	// 获取字段数量
	fmt.Println("numfiled :",d.NumField())

	//获取索引为2的字段
	floatfiled := d.Field(2)

	// 输出字段类型
	fmt.Println("Filed :", floatfiled.Type())

	// 根据名字查找字段
	fmt.Println("FiledByName(\"b\") .Type :", d.FieldByName("b").Type())

	// 根据索引查找值， next字段的int字段的值
	// 获取结构体dummy下标为4的字段 （next ,是一个结构体） 再找第四个字段的第一个字段  相当于切片
	fmt.Println("FiledByIndex([]int{4,0}).Type :", d.FieldByIndex([]int{4,0}).Type())
}