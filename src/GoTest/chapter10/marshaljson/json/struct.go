package json

import (
	"bytes"
	"reflect"
)

func WriteStruct(buff *bytes.Buffer, value reflect.Value) error  {
	//取值的类型对象
	valueType := value.Type()

	buff.WriteString("{")

	for i:=0;i<value.NumField();i++{
		//获取每个字段的值
		filedVaue := value.Field(i)

		// 获取每个字段的类型
		filedType := valueType.Field(i)

		// 写入字段名左双引号
		buff.WriteString("\"")

		//写入字段名
		buff.WriteString(filedType.Name)

		//写入字段名右双引号，和冒号
		buff.WriteString("\":")

		//写入每个字段值
		WriteAny(buff, filedVaue)

		// 每个字段尾部写入逗号，最后一个字段不添加
		for i <value.NumField()-1{
			buff.WriteString(",")
		}
	}

	buff.WriteString("}")

	return nil
}
