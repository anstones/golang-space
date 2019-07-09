package main

import "fmt"

func main()  {

	for x:= 0;x < 10; x ++{
		for y:= 0; y <10; y++{
			if y == 2{
				goto breakHere  // 推出多层循环
			}
		}
		return
	}
	breakHere:
		fmt.Println("done")
}

func noGoTo()  {

	var breakAgain bool

	for x:= 0;x<10; x++{
		for y:= 0;y<10;y++{
			if y == 2{
				breakAgain = true
				break
			}
		}
		if breakAgain{
			break
		}
	}
	fmt.Println("done with no goto")
}


