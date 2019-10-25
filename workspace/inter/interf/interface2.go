package interf

import "fmt"

type Test interface {
	Tester()
}

type Myfloat float64


func (m Myfloat) Tester()  {
	fmt.Println(m)
}

func Describe(T Test)  {
	fmt.Printf("Interface type %T value %v\n", T, T)
}

func Assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}