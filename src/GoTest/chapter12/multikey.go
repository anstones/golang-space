package main

// 利用map特性多键索引及查询
// GO 底层会为map的健自动构建哈希值

import (
	"fmt"
)

type Profile struct {
	Name string
	Age int
	Married bool
}

type queryKey struct {
	Name string
	Age int
}

var mapper = make(map[queryKey]*Profile)

func buildIndex(list []*Profile)  {

	for _, profile := range list{
		key := queryKey{
			Name:profile.Name,
			Age:profile.Age,
		}
		mapper[key] = profile
	}
}

func queryData(name string, age int)  {
	key := queryKey{name,age}
	result,ok := mapper[key]
	if ok{
		fmt.Println(result)
	}else {
		fmt.Println("not found")
	}
}

func main() {
	list := []*Profile{
		{"张三",36,true},
		{"李四",30,false},
		{"王五",20,true},
	}

	buildIndex(list)

	queryData("张三",36)
}
