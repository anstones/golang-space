package main

import (
	"fmt"
)

func main() {
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])

	fmt.Println("personSalary map contents:", personSalary)

	personSalary["mike"] = 9000
	newEmp := "joe"
	value, ok := personSalary[newEmp]  // 判断map是否包含元素
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp,"not found")
	}

	fmt.Println("All items of a map")
	for key, value := range personSalary {   //遍历map,每次遍历出的顺序不一定相同
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	delete(personSalary, "steve")  //删除map元素
	fmt.Println("after delete personSalary :", personSalary)

	personSalary["stone"] = 20000
	fmt.Println("Original person salary", personSalary)
	newpersonSalary := personSalary  //map是引用类型
	newpersonSalary["stone"] = 25000
	fmt.Println("Original person salary", personSalary)



}