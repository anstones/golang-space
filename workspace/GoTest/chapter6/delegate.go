package main

import "fmt"

type Class struct {
}

func (c *Class) Do(v int) {
	fmt.Println("call method do : ", v)
}

func Do(v int) {
	fmt.Println("call function do : ", v)
}

func main() {
	var delegate func(int)

	c := new(Class)

	delegate = c.Do

	delegate(100)

	delegate = Do

	delegate(100)
}
