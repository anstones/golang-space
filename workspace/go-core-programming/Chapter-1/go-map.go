package main

import "fmt"

// map 支持的操作:
//map 的单个键值访问格式为mapName[key］,更新某个key 的值时mapName[key］放到等号左边，访问某个key 的值时mapName[key］放在等号的右边。
//可以使用range 遍历一个map 类型变量，但是不保证每次选代元素的顺序。
//删除map 中的某个键值,使用如下语法:delete(mapName,key）.delete 是内置函数，用来删除map 中的某个键值对。
//可以使用内置的len()函数返回map 中的键值对数量。

func main() {
	mp := map[string]int{"a": 1, "b": 2}
	fmt.Println(mp["a"])
	fmt.Println(mp["b"])

	m1 := make(map[int]string)     //map 的容量使用默认位
	m2 := make(map[int]string, 10) //map 的容量使用给定的len 值
	m1[1] = "tom"
	m2[2] = "pony"
	fmt.Println(m1[1])
	fmt.Println(m2[2])

	m := make(map[int]string)
	m[1] = "tom"
	m[1] = "pony"
	m[2] = "jaky"
	m[3] = "andes"
	delete(m, 3)
	fmt.Println(m[1])
	fmt.Println(len(m))   //len 函数返回map 中的键值对的数量
	for k, v := range m { //range 支持边历mp ，但不保证每次边历次序是一样的
		fmt.Println("key=", k, "value=", v)
	}
}
