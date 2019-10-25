package main

import "fmt"

const MAX int = 3

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var a int = 20 /* 声明实际变量 */
	//var ip *int     /* 声明指针变量 */

	ip := &a

	fmt.Printf("a 的值：%d\n", a)

	fmt.Printf("a 的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 的值: %d\n", *ip)

	// 指针数组
	b := []int{10, 20, 30}
	var ptr [MAX]*int

	for i := 0; i < len(b); i++ {
		ptr[i] = &b[i]
		fmt.Printf("b[%d]的地址：%x\n", i, ptr[i])
	}

	for i := 0; i < len(b); i++ {
		fmt.Printf("b[%d]值：%d\n", i, *ptr[i])
	}

	//指针传递给函数
	x, y := 100, 200
	fmt.Printf("交换前x的值：%d\n", x)
	fmt.Printf("交换前y的值：%d\n", y)

	// 把x,y的地址传递给函数，然后x,y的地址交换，最终实现两个值的交换
	swap(&x, &y)

	fmt.Printf("交换后x的值：%d\n", x)
	fmt.Printf("交换后y的值：%d\n", y)

}
