package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	inputer := bufio.NewReader(os.Stdin)
	fmt.Println("please impurt you name")
	input , err := inputer.ReadString('\n') // 注意是单引号
	if err != nil{
		fmt.Printf("Fond an error: %s\n" ,err )
	}else {
		input = input[:len(input)-1]
		fmt.Printf("hello, %s\n", input)
	}
}