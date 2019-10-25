package main

// 类型的预测和转换

import "fmt"

func calssifier(items ...interface{})  {
	for i,x := range items{
		switch x.(type) {
		case bool:
			fmt.Printf("参数 #%d 的类型是 bool \n", i)
		case float64:
			fmt.Printf("参数 #%d 的类型是 float64 \n", i)
		case int, int64:
			fmt.Printf("参数 #%d 的类型是 int \n", i)
		case nil:
			fmt.Printf("参数 #%d 的类型是 nil \n", i)
		case string:
			fmt.Printf("参数 #%d 的类型是 string \n", i)
		default:
			fmt.Printf("参数 #%d 为未知类型\n", i)
		}

	}
}

func main() {
	calssifier(13,-14.3,"BELGIUM",complex(1,2), nil,false)
}
