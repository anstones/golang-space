package chapter3

import (
	"fmt"
	"sort"
)

func main()  {
	scene := make(map[string] int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	var sceneList []string

	for k,_ := range scene{
		sceneList = append(sceneList, k)
	}

	sort.Strings(sceneList)  // 切片排序
	fmt.Println(sceneList)
}
