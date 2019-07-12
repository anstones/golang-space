package main

import (
	"fmt"
	"reflect"
)

func main() {
	type cat struct {

	}

	// 实例化
	ins := &cat{}

	typeOfCat := reflect.TypeOf(ins)

	fmt.Printf("name: %v kind: %v\n",typeOfCat.Name(), typeOfCat.Kind() )

	Cat := typeOfCat.Elem()


	fmt.Printf("element name: %v, element kind: %v\n", Cat.Name(), Cat.Kind())


}