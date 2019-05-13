package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human) SayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string)  {
	fmt.Printf("La la la la...", lyrics)
}

func (e Employee) SayHi()  {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main()  {
	mike := Student{Human{"mike",25,"222-222-xxx"},"MIT",0.00}
	pol := Student{Human{"pol",30,"111-222-xxx"},"HAR",100}
	sam := Employee{Human{"sam",35,"333-222-xxx"},"Goland Inc",200}
	Tom := Employee{Human{"Tom",34,"444-222-xxx"},"Things Ltd",4000}

	var i Men
	i = mike
	fmt.Println("this is mike, a student:")
	i.SayHi()
	i.Sing("November rain")

	i = Tom
	fmt.Println("this is mike, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)    //T这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = pol, sam, mike
	for _, value := range x{
		value.SayHi()
	}
}

