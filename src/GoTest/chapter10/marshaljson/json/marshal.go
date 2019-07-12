package json

import (
	"bytes"
	"reflect"
)

func MarshalJson(v interface{})(string, error)  {
	var b bytes.Buffer

	if err := WriteAny(&b, reflect.ValueOf(v)); err == nil{
		return b.String(), nil
	}else {
		return "", nil
	}
}
