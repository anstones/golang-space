package main

import "fmt"

func main() {
	var any interface{}
	any = 1 // 空接口可以保存任意类型值
	fmt.Println(any)

	var a int = 1
	var i interface{} = a
	var b int = i.(int) //从空接口中取值, 需要用类型断言（因为空接口可以存储任意类型数据）
	fmt.Println(b)

	var x interface{} = 100
	var y interface{} = "100"
	fmt.Println(x == y) //不同类型的空接口 比较的结果不同

	var c interface{} = []int{10}
	var d interface{} = []int{10}
	fmt.Println(c == d) // 不能比较空接口中的动态值

}
