package interf

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	EmpId    int
	Basicpay int
	Pf       int
}

type Contract struct {
	EmpId  int
	Basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.Basicpay + p.Pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.Basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
*/
func TotalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		fmt.Println(v)
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

