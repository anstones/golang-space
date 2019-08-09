package main

import (
	"fmt"
	"strings"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

type Employee struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

// 形式传参中传递的是一个副本，函数的中操作的也是副本，操作不会影响结构体
func operateEmployee1(employee Employee) {
	employee.ID = 10010
}

// 指针传参中传递的是引用地址，函数的中操作会影响结构体
func operateEmployee2(employee *Employee) {
	employee.ID = 10010
}

func main() {
	var pers1 Person
	pers1.firstName = "张"
	pers1.lastName = "三"
	upPerson(&pers1)
	fmt.Printf("这个人的名字是：%s%s\n", pers1.firstName, pers1.lastName)

	pers2 := new(Person)
	pers2.firstName = "张"
	pers2.lastName = "三"
	upPerson(pers2)
	fmt.Printf("这个人的名字是：%s%s\n", pers2.firstName, pers2.lastName)

	pers3 := &Person{"张", "三"}
	upPerson(pers3)
	fmt.Printf("这个人的名字是：%s%s\n", pers3.firstName, pers3.lastName)

	var employee Employee
	employee.ID = 1001
	employee.Name = "Tom"
	employee.Address = "xxxxx"
	employee.Phone = "18814128405"

	fmt.Printf("形式传参之前，employee ID: %d\n", employee.ID)
	operateEmployee1(employee)
	fmt.Printf("形式传参之后，employee ID: %d\n", employee.ID)

	fmt.Printf("指针传参之前，employee ID: %d\n", employee.ID)
	operateEmployee2(&employee)
	fmt.Printf("指针传参之后，employee ID: %d\n", employee.ID)

	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "less 2")
	}

}
