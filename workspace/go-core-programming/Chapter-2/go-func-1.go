package main

import "fmt"

func Chvalue(a int) int {
	a = a + 1
	return a
}

func Chpointer(a *int) {
	*a = *a + 1
	return
}

func Sum(arr ...int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func main() {
	a := 10
	Chvalue(a) //实参传递给形参是值拷贝
	fmt.Println(a)

	Chpointer(&a) //实参传递给形参仍然是值拷贝，只不过复制的是a 的地址位
	fmt.Println(a)

	silce := []int{1, 2, 3, 4}
	array := [...]int{1, 2, 3, 4}
	// Sum(array)     //数纽不可以作为实参传递给不定参数的函数
	// Sum(array...)  //数纽不可以作为实参传递给不定参数的函数
	fmt.Println(Sum(silce...))
}
