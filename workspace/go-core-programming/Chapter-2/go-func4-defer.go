package main

import "fmt"

//Go 函数里提供了defer 关键字，可以注册多个延迟调用，这些调用以先进后出（ FILO ）的
//顺序在函数返回前被执行。这有点类似于Java 语言中异常处理中的finaly 子句。defer 常用于保
//证一些资源最终一定能够得到回收和释放。

// defer 后面必须是函数或方法的调用，不能是语句


func main() {
	defer func() {
		fmt.Println("first")
	}()

	defer func() {
		fmt.Println("second")
	}()

	fmt.Println("function body")


}
