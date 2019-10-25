package main

import (
	"fmt"
)

type Rectangle struct {
	width float64
	height float64
}

func area(r Rectangle) float64  {
	return r.width * r.height
}

func main()  {
	r := Rectangle{1,2}
	r2 := Rectangle{1,10}
	fmt.Printf("Area of r1 is: \n", area(r))
	fmt.Printf("Area of r1 is: \n", area(r2))
}
