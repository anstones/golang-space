package json

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
)

func WriteAny(buff *bytes.Buffer, value reflect.Value) error  {
	switch value.Kind() {
	case reflect.String:
		// 写入带有双引号括起来的字符串
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return WeiteSlice(buff,value)
	case reflect.Struct:
		return WriteStruct(buff,value)
	default:
		// 遇到不认识的种类，返回错误
		return errors.New("unsupport kind:" +value.Kind().String())

	}

	return nil
}
