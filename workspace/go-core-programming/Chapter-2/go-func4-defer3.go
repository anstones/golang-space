package main

import (
	"fmt"
	"os"
)

//主动调用as.Exit( int） 退出进程时， defer 将不再被执行（即使defer 已经提前注册〉。

func main() {
	defer func() {
		fmt.Println("first")
	}()
	fmt.Println("func body")

	os.Exit(1)
}
