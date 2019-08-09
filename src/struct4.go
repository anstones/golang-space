package main

import "fmt"

type File struct {
	fd   int
	name string
}

// 工厂方法创建结构体实例
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main() {
	f := NewFile(10, "/test.txt")

	fmt.Println(f)
}
