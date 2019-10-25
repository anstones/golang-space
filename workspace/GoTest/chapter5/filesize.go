package main

import (
	"fmt"
	"os"
)

func fileSize(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil{
		return 0
	}
	defer f.Close()

	info, err := f.Stat()

	if err != nil{
		// defer机制触发
		return 0
	}

	size := info.Size()

	// defer 机制触发
	return size
}


func main()  {
	fmt.Println(fileSize("C:\\Users\\Administrator\\Desktop\\oeasy\\data\\mine\\Project\\Go_Space\\src\\GoTest\\chapter5\\accumulator.go"))
}
