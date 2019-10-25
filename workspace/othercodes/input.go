package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input you name:")
	input,err := inputReader.ReadString('\n')
	if err != nil{
		fmt.Printf("found an error: %s\n", err)
	}else{
		input = input[:len(input)-1]
		fmt.Printf("hello, %s\n", input)
	}
}
