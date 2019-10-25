package main

import (
	"GoTest/chapter8/clsfactory/base"
	_ "GoTest/chapter8/clsfactory/cls1"
	_ "GoTest/chapter8/clsfactory/cls2"
)

func main()  {
	c1 := base.Create("Class1")
	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()


}