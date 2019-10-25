package main

// 可变参数的传递
import "fmt"

func rawPrint(rawList ...interface{})  {
	for _, s := range rawList{
		fmt.Println(s)
	}
}

func Print(slist ...interface{})  {  //slist 类型为[]interface{}

	// 经slist可变参数的切片完整的传递给下一个函数
	rawPrint(slist...)
}

func main()  {
	Print(1,2,3)
}