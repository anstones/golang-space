package main

import (
	"fmt"
	"os"
)


func main() {
	f, err := os.Create("C:\\Users\\Administrator\\Desktop\\oeasy\\data\\mine\\Project\\Go_Space\\src\\writefile\\test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	b := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}//将bytes 写入文件
	l, err := f.Write(b)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}