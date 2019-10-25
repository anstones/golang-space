package main

import "fmt"

func main() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)

}

func ap(a []int) {
	a = append(a, 10) // append 导致底层数组重新分配内存了，ap 中的 a 这个slice 的底层数组和外面的不是一个，并没有改变外面的。
}

func app(a []int) {
	a[0] = 1
}

type MyInt1 int   // 基于类型 int 创建了新类型 MyInt1
type MyInt2 = int // 创建了 int 的类型别名 MyInt2，注意类型别名的定义是 =

func main2() {
	var i int = 0
	//var i1 MyInt1 = i  // 将 int 类型的变量赋值给 MyInt1 类型的变量， 编译不通过
	var i1 MyInt1 = MyInt1(i) // 使用强制类型转化

	var i2 MyInt2 = i // MyInt2 只是 int 的别名，本质上还是 int，可以赋值
	fmt.Println(i1, i2)
}
