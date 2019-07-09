package chapter2

import "fmt"

func main()  {

	// 变量，指针， 地址 三者之间的关系： 每个变量都拥有地址， 指针的值就是地址

	var ponit =
		`
		取地址操作符"&"和取值操作符"*"是一对互补操作符,"&"取出地址，"*"根据地址取出地址指向的值
		变量，指针地址，指针变量，取地址，取值的相互关系和特性：
		1： 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
		2： 指针变量的值是指针地址
		3： 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值
		`
	fmt.Printf("%s", ponit)

	var house = "Malibu Point 10880, 90265"

	// 对house取地址
	ptr := &house

	// 打印ptr类型
	fmt.Printf("ptr type :%T\n", ptr)

	// 打印ptr地址
	fmt.Printf("address： %p\n", ptr)

	// 对指针取值
	vaule := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n",vaule)

	// 指针取值后就是指向变量的值
	fmt.Printf("value :%s\n",vaule)

	x, y := 1,2

	swap(&x,&y)

	fmt.Println(x,y)

	swap2(x,y)

	fmt.Println(x,y)
}

func swap(a,b *int)  {
	// 取a指针的值，赋值给t
	t := *a

	// 取b指针的值，赋值给a指针指向的变量
	*a = *b

	// 将a指针的值赋值给b指针指向的变量
	*b = t

}

func swap2(c, d int)  {
	c,d = d,c
}