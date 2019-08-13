package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newfile  *os.File
	fileInfo os.FileInfo
	err      error
	path     = "test/test2/"
	fileName = "demo"
	filePath = path + fileName
)

func main() {
	// 创建文件夹
	err := os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Println("创建目录成功")
	}

	//创建文件
	newfile, err = os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newfile)
	newfile.Close()

	// 查看文件信息
	fileInfo, err = os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("===============")
		log.Fatal("文件不存在")
	}

	fmt.Println("文件名：", fileInfo.Name())
	fmt.Println("文件大小：", fileInfo.Size())
	fmt.Println("文件权限：", fileInfo.Mode())
	fmt.Println("最后修改时间：", fileInfo.ModTime())
	fmt.Println("是否是文件夹：", fileInfo.IsDir())
	fmt.Printf("系统接口类型：%T\n", fileInfo.Sys())
	fmt.Printf("系统信息：%+v\n\n", fileInfo.Sys())

	newpath := filePath + ".txt"
	err = os.Rename(filePath, newpath) //移动或重命名文件/文件夹
	if err != nil {
		log.Fatal(err)
	}
}
