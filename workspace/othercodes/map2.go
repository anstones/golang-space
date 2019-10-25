package main

import (
	"fmt"
)

func init()  {
	fmt.Println("---------------比较两个map 是否相等-------------------------")
}

func main()  {
	mapOne := map[string]string{"uid":"1","nickname":"jeck","age":"28"}
	mapTwo := map[string]string{"uid":"1","nickname":"sss","age":"28"}
	is_identical := 0
	if len(mapOne) == len(mapTwo) {
		for key,_ := range mapOne {
			if mapTwo[key] != mapOne[key] {
				is_identical = 1
			}
		}
	}else{
		is_identical = 1
	}
	if is_identical > 0 {
		fmt.Println("mapOne和mapTwo不相等")
	}else{
		fmt.Println("mapOne和mapTwo相等")
	}

}

