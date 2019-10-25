package interf

import (
	"fmt"
)

type SalaryCalculators interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	FirstName string
	LastName string
	BasicPay int
	Pf int
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.FirstName, e.LastName, (e.BasicPay + e.Pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.TotalLeaves - e.LeavesTaken
}


