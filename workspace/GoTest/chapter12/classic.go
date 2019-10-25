package main

// 基于哈希值的多键索引及查询

import "fmt"

type Profile struct {
	Name string
	Age int
	Married bool
}

// 实例化map  健为整形类型，保存哈希值；值类型为*profile，
var mapper = make(map[int][]*Profile)

// 申明查询健结构，包含name和age
type classicQuerykey struct {
	Name string
	Age int
}

// 计算查询健的哈希值
func (c *classicQuerykey) hash() int {
	//将名字的哈希和年龄的哈希相加
	return simpleHash(c.Name) + c.Age*1000000
}


// 构建数据索引
func buildIndex(list []*Profile)  {
	//遍历所有元素
	for _, profile := range list{
		// 计算查询健的哈希值， 将名字和年龄赋值到查询健中进行保存
		key := classicQuerykey{profile.Name, profile.Age}
		// 计算数据的哈希值，取出已存在的记录
		existValue := mapper[key.hash()]
		// 将当前数据添加到已存在的记录切片中
		existValue = append(existValue, profile)
		// 将切片重新设置到映射中
		mapper[key.hash()] = existValue
	}
}

func queryData(name string, age int)  {
	// 根据指定查询条件查询值
	keyToQuery := classicQuerykey{name,age}
	// 计算查询健的哈希值并查询，获得相同哈希值的所有结果集合
	resultList := mapper[keyToQuery.hash()]
	// 遍历结果集
	for _,result := range resultList{
		// 与查询结果比对
		if result.Name == name && result.Age == age{
			fmt.Println(result)
			return
		}
	}
	fmt.Println("not found")
}

// 字符串转哈希
func simpleHash(str string) (ret int)  {
	// 遍历字符串中的每个字符（ASCII）
	for i :=0 ; i<len(str);i++{
		// 取出字符
		c:= str[i]
		//将字符的ASCII相加
		ret += int(c)
	}
	return ret
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
