package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type NamePoint struct {
	Point
	name string
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *NamePoint) Abs() float64 {
	return p.Point.Abs() * 100
}

func main() {
	n := &NamePoint{Point{3, 4}, "python"}

	// Point的Abs方法晋升为外层类型的方法
	fmt.Println(n.Abs())

	fmt.Println(n.Abs())
}
