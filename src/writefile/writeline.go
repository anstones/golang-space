package main

import (
	"fmt"
	"os"
)

func main()  {
	f, err := os.Create("C:\\Users\\Administrator\\Desktop\\oeasy\\data\\mine\\Project\\Go_Space\\src\\writefile\\test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	b := []string{"Welcome to the world of Go1.", "Go is a compiled language.","It is easy to learn Go."}//将string 逐行写入文件
	for _, v := range b{
		fmt.Fprintln(f,v)   //逐行写入文件
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}

	fmt.Println("file written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}