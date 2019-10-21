package main

import "fmt"

//1.关于 cap() 函数的适用类型，下面说法正确的是()
//A. array
//B. slice
//C. map  // cap() 函数不适用 map。
//D. channel

func main() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")

	s := make(map[string]int)
	delete(s, "h")      // 删除 map 不存在的键值对时，不会报错，相当于没有任何作用；
	fmt.Println(s["h"]) // 获取不存在的键值对时，返回值类型对应的零值，所以返回 0
}
