package main

import (
	"log"
	"os"
)

var (
	newfile  *os.File
	fileInfo os.FileInfo
	err      error
	path     = "test/test2/"
	fileName = "demo.txt"
	filePath = path + fileName
)

func main() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	file.Close()

	// 文件路径， 打开时的属性，打开时的文件权限模式
	file, err = os.OpenFile(path, os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
