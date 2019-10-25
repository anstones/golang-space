package main

import "fmt"

func playerGen(name string) func() (string, int){

	hp := 150  // 外部不能直接访问hp
	return func() (string,int) {
		return name, hp
	}
}

func main()  {
	generator := playerGen("high noon")

	name, hp := generator()

	fmt.Println(name, hp)
}