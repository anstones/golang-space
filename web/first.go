package main

import "fmt"

type Person struct {
	name string
	age int
}

func Oder(p1, p2 Person)(Person, int)  {
	if p1.age > p2.age{
		return p1, p1.age-p2.age
	}
	return  p2, p2.age-p1.age
}

func main()  {
	var tom Person
	tom.name, tom.age = "tom", 18
	bob := Person{age:20, name:"James"}
	pol := Person{"pol",25}

	a,b := Oder(tom,bob)
	fmt.Printf("Of %s and %s, %s is older by %d years\n",        tom.name, bob.name, a.name, b)
	c,d := Oder(tom,pol)
	fmt.Printf("Of %s and %s, %s is older by %d years\n",        tom.name, pol.name, c.name, d)
}