package json

import (
	"bytes"
	"reflect"
)

func WeiteSlice(buff *bytes.Buffer, value reflect.Value) error  {
	// 写入切片开始标记
	buff.WriteString("[")

	//遍历每个元素
	for s:=0;s<value.Len();s++{
		sliceValue := value.Index(s)
		// 写入每个切片元素
		WriteAny(buff, sliceValue)

		if s < value.Len()-1{
			buff.WriteString(",")
		}
	}

	buff.WriteString("]")

	return nil
}