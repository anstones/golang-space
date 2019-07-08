package main

import (
	"fmt"
)


func append() func(s string)string{
	t := "hello"
	s := func(s string) string {
		t = t + " " + s
		return t //现在 t 值变为了 Hello World。
	}
	return s
}
func main()  {
	a := append()
	b := append()
	fmt.Println(a("world"))
	fmt.Println(b("Everyone"))
	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}
