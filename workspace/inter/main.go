package main

import "fmt"
import "inter/interf"
import "reflect"

func main() {
	// 第一例
	name := interf.MyString("Sam Anderson")
	fmt.Println(name)
	var v interf.VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Println("====", reflect.TypeOf(v))
	fmt.Printf("Vowels are %c", v.FindVowels())

	// 第二例
	fmt.Printf("\n")
	pemp1 := interf.Permanent{1, 5000, 20}
	pemp2 := interf.Permanent{2, 6000, 30}
	cemp1 := interf.Contract{3, 3000}
	employees := []interf.SalaryCalculator{pemp1, pemp2, cemp1}
	interf.TotalExpense(employees)


	// 第三例
	fmt.Printf("\n")
	var t interf.Test
	f := interf.Myfloat(89.7)
	t = f
	interf.Describe(t)
	t.Tester()

	// 第四例 空接口用作断言
	var s interface{} = 56
	interf.Assert(s)
	var i interface{} = "hello world"
	interf.Assert(i)

	// 第五例 空接口用作类型选择
	interf.FindType("hello world")
	p := interf.Person{
		Name:"james",
		Age: 34,
	}
	interf.FindType(p)

	// 第六例 指针接受者与值接受者
	p1 := interf.Personer{"AD",26}
	var d1 interf.Describtion = p1
	d1.Describer()

	p2 := interf.Personer{"AK",40}
	d1 = &p2
	d1.Describer()

	a := interf.Address{"beijing","china"}
	var d2 interf.Describtion = &a //Address 类型的指针实现了 Describer 接口
	d2.Describer()

	//第七例 多个接口
	e := interf.Employee {
		FirstName: "Naveen",
		LastName: "Ramanathan",
		BasicPay: 5000,
		Pf: 200,
		TotalLeaves: 30,
		LeavesTaken: 5,
	}
	fmt.Println(e)
	var ss interf.SalaryCalculators = e
	ss.DisplaySalary()
	var ll interf.LeaveCalculator = e
	fmt.Printf("\nLeaves left = %d", ll.CalculateLeavesLeft())

}