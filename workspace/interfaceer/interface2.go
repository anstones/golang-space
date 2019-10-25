package main

import (
	"fmt"
	"math"
)

type Square struct {

	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32{
	return sq.side*sq.side
}

func (c *Circle) Area() float32  {
	return c.radius*c.radius*math.Pi
}

func main()  {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1

	if t, ok := areaIntf.(*Square); ok{
		fmt.Printf("areaIntf 的类型是：%T\n",t)
	}

	if u, ok := areaIntf.(*Circle);ok{
		fmt.Printf("areaIntf 的类型是：%T\n",u)
	}else{
		fmt.Printf("areaIntf 不含类型为Circle的变量")
	}

}