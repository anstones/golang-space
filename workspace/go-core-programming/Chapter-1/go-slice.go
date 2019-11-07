package main

import "fmt"

// 切片支持的操作:
// 内置函数len()返回切片长度。
// 内置函数cap()返回切片底层数组容量。
// 内置函数append() 对切片追加元素。
// 内置函数copy()用于复制一个切片。

func main() {
	var array = [...]int{0, 1, 2, 3, 4, 5, 6} //创建有7 个int 型元素的数纽
	s1 := array[0:4]
	s2 := array[:4]
	s3 := array[2:]
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)

	a := make([]int, 10)     //len=lO , cap=lO
	b := make([]int, 10, 15) //len=lO , cap=l5
	fmt.Printf("%v\n", a)    //[0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("%v\n", b)    //[0 0 0 0 0 0 0 0 0 0]

	x := [...]int{0, 1, 2, 3, 4, 5, 6}
	y := make([]int, 2, 4)
	z := x[0:3]

	fmt.Println(len(y)) //2
	fmt.Println(cap(y)) //4
	y = append(y, 1)
	fmt.Println(y)      //[0,0, 1]
	fmt.Println(len(y)) //3

	y = append(y, z...)
	fmt.Println(y)      //[0 0 1 0 1 2)
	fmt.Println(len(y)) //6
	fmt.Println(cap(y)) //cap(y)=8，底层数纽发生扩展

	d := make([]int, 2, 2)
	copy(d, z)          // //copy 只会主制d 和c 中长度最小的
	fmt.Println(d)      //[0,1]
	fmt.Println(len(d)) //2
	fmt.Println(cap(d)) //2

	str := "hello ， 世界！" //通过字符串字面量初始化一个字符串str
	m := []byte(str)     //将字符串转换为[]byte 类型切片
	n := []rune(str)     //将字符串转换为[]rune 类型切片

}
