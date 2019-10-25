package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestString := "Hi, Golang"
	MD5Inst := md5.New()
	MD5Inst.Write([]byte(TestString))
	result := MD5Inst.Sum([]byte(""))
	fmt.Printf("%x\n", result)

	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(TestString))
	Result := sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n", Result)

}
