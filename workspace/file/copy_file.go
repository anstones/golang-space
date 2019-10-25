package main

import (
	"io"
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
	// 打开文件
	oldFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer oldFile.Close()

	// 创建新的文件作为目标文件
	newFile, err := os.Create(filePath + "_copy")
	if err != nil {
		log.Fatal(err)

	}
	defer newFile.Close() //create函数执行后需要close关闭回收资源

	//从原文件复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, oldFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("文件已复制，大小 %d bytes。", bytesWritten)

	// 将文件内容flush到硬盘中
	err = newFile.Sync() // 调用io.copy之后文件内容并没有正真保存在文件中，需要使用sync函数同步之后才能正真保存到硬盘中
	if err != nil {
		log.Fatal(err)
	}

}
