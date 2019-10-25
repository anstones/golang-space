package main

import(
	"fmt"
	"math"
)
type Rectangle struct {
	width float64
	height float64
}

type Circle struct {    
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}
func (c Circle) area() float64 {
	return c.radius*c.radius*math.Pi
}

func main()  {
	r1 := Rectangle{8,9}
	c1 := Circle{20}
	fmt.Printf("Area of r1: ", r1.area())
	fmt.Printf("Area of r1: ",c1.area())
}