// 多态
package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours  int
	hourlyRate int
}

type Add struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func (a Add)calculate() int{
	return a.noOfClicks * a.CPC
}

func (a Add)source() string {
	return a.adName
}

func calculateNetIncome(ic []Income){ // 接口类型的切片作为参数
	var netincome int = 0
	for _,v := range ic{
		fmt.Printf("Income From %s = $%d\n", v.source(), v.calculate())
		netincome += v.calculate()
	}
	fmt.Printf("Income From  = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}

	bannerAd := Add{adName: "Banner Ad", CPC: 2, noOfClicks: 500} // 新增收益流，但却完全没有修改calculateNetIncome函数
	popupAd := Add{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}  // 创建了Income类型的切片，存放三个项目
	calculateNetIncome(incomeStreams)
}