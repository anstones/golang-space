package main

import "fmt"

func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

type Person struct {
	age int
}

func main() {
	f()
	fmt.Println("M")

	person := &Person{28}

	defer fmt.Println(person.age) //将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28

	defer func(p *Person) {
		fmt.Println(p.age) //处.defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后面的函数执行的时候取出仍是 28
	}(person)

	defer func() {
		fmt.Println(person.age) //闭包引用，person 的值已经被改变，指向结构体 Person{29}，所以输出 29.
	}()

	person = &Person{29} //修改引用对象本身
}
