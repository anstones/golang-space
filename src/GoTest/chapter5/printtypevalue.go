package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string { //可变参数为interface{}  可以传入任何类型的值
	var b bytes.Buffer

	for _,s := range slist{
		str := fmt.Sprintf("%v",s)

		var typeString string

		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		default:
			typeString = "unknow type"
		}
		b.WriteString("value: ")
		b.WriteString(str)
		b.WriteString(" type: ")
		b.WriteString(typeString)
		b.WriteString("\n")

	}
	return b.String()
}

func main()  {
	fmt.Println(printTypeValue(100,"str",true))
}