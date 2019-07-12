package main

// 反射对象的空值和有效性判断

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int
	fmt.Println("var s *int :",reflect.ValueOf(a).IsNil())

	fmt.Println("nil: ", reflect.ValueOf(nil).IsValid())

	// *int 类型的空指针
	fmt.Println("(*int)(nil):",reflect.ValueOf((*int)(nil)).Elem().IsValid())

	s := struct {

	}{}

	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员：",reflect.ValueOf(s).FieldByName("").IsValid())

	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法：", reflect.ValueOf(s).MethodByName("").IsValid())


}

