package main

import "io"

// 接口同结构体一样，也可以嵌套，嵌套后组成新接口，只要实现了接口的所有方法，则这个新接口中的所有嵌套接口的方法均可被调用

type Device struct {
}

func (d *Device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *Device) Close() error {
	return nil
}

func main() {
	var wc io.WriteCloser = new(Device) //wc  为io.WriteCloser 接口 嵌套了Write和Close接口

	wc.Write(nil)

	wc.Close()

	var writeOnly io.Writer = new(Device)
	writeOnly.Write(nil)
}
