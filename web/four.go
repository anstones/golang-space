package main

import (
	"fmt"
)

type Humen struct {
	name string
	age int
	phone string
}

type Employee struct {
	Humen
	speciality string
	phone string
}

func main()  {
	Bob := Employee{Humen{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)    // 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Humen.phone)
}
