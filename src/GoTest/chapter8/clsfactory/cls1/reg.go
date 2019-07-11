package cls1

import (
	"GoTest/chapter8/clsfactory/base"
	"fmt"
)

type Class1 struct {

}

func (c *Class1) Do()  {
	fmt.Println("Classs1")
}

func init()  {
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})
}
