package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改值
// 值可修改的条件： 1：可被寻址； 2：可被导出（字段首字母大写）

func main() {
	var a int = 1024

	type dog struct {
		//legCount int //不可被导出
		LegCount int
	}

	//valueOfA := reflect.ValueOf(a) // 传入的时a的值，而不是地址，不能被寻址,所以修改值会报错
	//valueOfA.SetInt(1)  // 报错


	valueOfA := reflect.ValueOf(&a)   // 获取变量a的反射对象（a的地址）

	valueOfA = valueOfA.Elem()  // 取出a的地址的元素（a指向的值）

	valueOfA.SetInt(1)  // 修改a的值

	fmt.Println(valueOfA.Int())


	//valueOfDog := reflect.ValueOf(dog{})
	//vLegCount := valueOfDog.MethodByName("legCount")   //legCount 不可被导出
	//vLegCount.SetInt(4)

	valueOfDog := reflect.ValueOf(&dog{})
	valueOfDog = valueOfDog.Elem()
	vLegCount := valueOfDog.FieldByName("LegCount")
	vLegCount.SetInt(4)
	fmt.Println(vLegCount.Int())
}
