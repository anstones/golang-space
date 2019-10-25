package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	newfile  *os.File
	fileInfo os.FileInfo
	path     = "test/test2/"
	fileName = "demo.txt"
	filePath = path + fileName
)

func main() {
	// 使用linux风格改变文件权限
	err := os.Chmod(filePath, 0777)
	if err != nil {
		log.Panicln(err)
	}

	// 改变文件的所有者
	err = os.Chown(filePath, os.Geteuid(), os.Getegid())
	if err != nil {
		log.Println(err)
	}

	//查看文件信息
	fileInfo, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("文件不存在")
		}
		log.Fatal(err)
	}

	fmt.Println("最后修改时间时：", fileInfo.ModTime())

	//改变时间戳
	twoDayFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDayFromNow
	lastModifyTime := twoDayFromNow
	err = os.Chtimes(filePath, lastAccessTime, lastModifyTime)
	if err != nil {
		log.Panicln(err)
	}
}
