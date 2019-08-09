package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {

	sq1 := new(Square)
	sq1.side = 5

	sq2 := new(Rectangle)
	sq2.length = 4
	sq2.width = 3

	a := []Shaper{sq1, sq2}

	for n, _ := range a {
		fmt.Println("参数是：", a[n])
		fmt.Println("面积是：", a[n].Area())
	}

}
