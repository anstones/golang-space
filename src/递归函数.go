package main

import (
	"fmt"
)

func main() {
	a := jiechen(10)
	fmt.Println(a)

	var i int
	for i=0; i<10;i++{
		b := feibo(i)
		fmt.Println(b)
	}
}

func jiechen(n int) (result int){
	if(n>0){
		result = n*jiechen(n-1)
		return result
	}
	return 1
}

func feibo(x int) (result int) {
	if x<2 {
		return x
	}
	return feibo(x-2) + feibo(x-1) 
}