package interf

import "fmt"

type Describtion interface {
	Describer()
}
type Personer struct {
	Name string
	Age  int
}

func (p Personer) Describer() { // 使用值接受者实现
	fmt.Printf("%s is %d years old\n", p.Name, p.Age)
}

type Address struct {
	State   string
	Country string
}

func (a *Address) Describer() { // 使用指针接受者实现
	fmt.Printf("State %s Country %s", a.State, a.Country)
}