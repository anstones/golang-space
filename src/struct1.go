package main

import (
	"fmt"
)

type Person struct {
	string  // 匿名字段
	int
}

func main() {

	emp1 := struct {   // 匿名结构体
		firstName, lastName  string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee ", emp1)

	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)
}