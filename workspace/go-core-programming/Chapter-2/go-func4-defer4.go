package main

//使用defer 改写后，在打开资源无报错后直接调用defer 关闭资源， 一旦养成这样的编程习惯，则很难会忘记资源的释放

import (
	"io"
	"os"
)

func CopyFile(dst, src string) (w int64, err error) {
	src, err := os.Open(src)
	if err != nil {
		return
	}
	dst, err := os.Create(dst)
	if err != nil {
		//src 很容易被忘记关闭
		src.Close()
		return
	}
	w, err := io.Copy(dst, src)

	dst.Close()
	src.Close()
	return
}

func CopyFileWithDefer(dst, src string) (w int64, err error) {
	src, err := os.Open(src)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dst.Close()
	w, err := io.Copy(dst, src)
	return
}

// defer 语句的位置不当，有可能导致panic ， 一般defer语句放在错误检查语句之后。
// defer 也有明显的副作用：
// defer 会推迟资源的释放，
// defer 尽量不要放到循环语句里面，将大函数内部的defer 语句单独拆分成一个小函数是一种很好的实践方式。
// defer 相对于普通的函数调用需要间接的数据结构的支持，相对于普通函数调用有一定的性能损耗。
// defer 中最好不要对有名返回值参数进行操作，否则会引发匪夷所思的结果

func main() {
	CopyFile("./file.txt", "./new_file.txt")
}
