package chapter3

import "fmt"

func main()  {

	//数组
	var team = [...]string{"h","s","m"}

	for k,v := range team{
		fmt.Println(k,v)
	}


}
