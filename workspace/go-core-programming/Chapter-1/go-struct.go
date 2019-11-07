package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	*Person
	Number int
}

func main() {
	P := &Person{
		Name: "Tom",
		Age:  20,
	}

	s := Student{
		Person: P,
		Number: 101,
	}

	fmt.Println(s.Person, s.Number)
}
