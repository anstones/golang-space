package main

import "fmt"
import "io/ioutil"




func checkerr(err error){
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
}

func main() {
	data,err := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\oeasy\\data\\mine\\Project\\Go_Space\\src\\filehanding\\test.txt") // 需要绝对路径
	checkerr(err)
	fmt.Println("Contents of file: ",string(data))
}