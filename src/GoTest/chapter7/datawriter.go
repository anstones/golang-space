package main

//接口的方法名 必须和 实现接口的类型方法名 完全一样
//实现接口的类型方法 的参数  返回值  必须和接口方法 的参数 返回值 一致

import "fmt"

type DateWriter interface {
	WriteData(data interface{}) error

	CanWrite() bool
}

type file struct {
}

func (d file) WriteData(data interface{}) error {
	fmt.Println("writedata: ", data)
	return nil
}

func (d file) CanWrite() bool { // 必须实现接口的所有方法
	return true
}

func main() {

	var writer DateWriter

	f := new(file)
	fmt.Printf("%T\n", f)
	writer = f
	writer.WriteData("data")
}
