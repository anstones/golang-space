package main

import (
	"fmt"
)

type Human struct {
	name, phone string
	age int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	compary string
}

func (h Human) sayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e Employee) sayHi()  {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,e.compary,e.phone)
}

func main()  {
	mark := Student{Human{"mark","123456",25},"MIT"}
	sam := Employee{Human{"sam","789654",42},"SMT"}
	mark.sayHi()
	sam.sayHi()
}