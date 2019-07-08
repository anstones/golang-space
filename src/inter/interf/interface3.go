package interf

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	Name string
	Age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.Name, p.Age)
}

func FindType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}
