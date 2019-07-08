// defer  延迟调用
package main

import "fmt"

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func main()  {
	a := 5
	defer printA(a)  // 实参取值 虽然延迟了，但是依然取的是 5
	a = 10
	fmt.Println("value of a before deferred function call", a)

	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)  //defer 栈 ,逆序
	}
}
