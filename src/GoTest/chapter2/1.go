package chapter2

import "fmt"

func sy()  {
	var a int = 100
	var b int = 200
	b, a = a, b
	fmt.Println(a,b)
}

func main()  {
	sy()
}
