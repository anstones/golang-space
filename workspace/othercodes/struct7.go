package main

import (
	"fmt"
)

type Camera struct {
}

type Phone struct {
}

type CameraPhone struct {
	Camera
	Phone
}

func (c *Camera) TakeAPicture() string {
	return "拍照"
}

func (p *Phone) Call() string {
	return "响铃"
}

func main() {
	cp := new(CameraPhone)
	// 多重嵌套，被继承的结构体的方法，晋升为继承结构体的方法，可以被调用
	fmt.Println("我打开相机：", cp.TakeAPicture())
	fmt.Println("电话来了：", cp.Call())
}
