package main

import (
	"fmt"
)

type valuable interface {
	getvalue() float32
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

type car struct {
	make  string
	model string
	price float32
}

func (s stockPosition) getvalue() float32 {
	return s.count * s.sharePrice
}

func (c car) getvalue() float32 {
	return c.price
}

func showvalue(asset valuable) {
	fmt.Printf("资产的价值是：%f\n", asset.getvalue())
}

func main() {
	var o valuable = stockPosition{"GooG", 577.20, 4}
	showvalue(o)
	o = car{"BMW", "M3", 66500}
	showvalue(o)
}
