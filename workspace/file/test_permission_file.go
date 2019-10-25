package main

import (
	"log"
	"os"
)

var (
	newfile  *os.File
	fileInfo os.FileInfo
	path     = "test/test2/"
	fileName = "demo.txt"
	filePath = path + fileName
)

func main() {
	// 测试写权限
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil && os.IsPermission(err) {
		log.Panicln("错误：没有写入权限。")
	} else if os.IsNotExist(err) {
		log.Panicln("错误： 文件不存在")
	} else {
		log.Fatal(err)
	}
	file.Close()

	//测试读权限
	file, err = os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil && os.IsPermission(err) {
		log.Println("错误：没有读取权限")
	} else if os.IsNotExist(err) {
		log.Println("错误：文件不存在")
	} else {
		log.Fatal(err)
	}

	file.Close()
}
