package main

import (
	"fmt"
	"structs/computer"
)



func main()  {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	//spec.model = "Mac Mini"  // model 为非导出字段
	fmt.Println("Spec:", spec)

}