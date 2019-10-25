package main

import (
	"fmt"
)


type Human struct {    
	name string    
	age int    
	phone string 
} 

type Student struct {    
	Human //匿名字段Human    
	school string    
	loan float32 
} 

type Employee struct {    
	Human //匿名字段Human    
	company string    
	money float32 
} 

//Human对象实现Sayhi方法 
func (h *Human) SayHi() {    
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone) 
} 

// Human对象实现Sing方法 
func (h *Human) Sing(lyrics string) {    
	fmt.Println("La la, la la la, la la la la la...", lyrics) 
} 

//Human对象实现Guzzle方法 
func (h *Human) Guzzle(beerStein string) {    
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein) 
}


// Employee重载Human的Sayhi方法 
func (e *Employee) SayHi() {    
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, 
	e.company, e.phone) //Yes you can split into 2 lines here. 
} 

//Student实现BorrowMoney方法 
func (s Student) BorrowMoney(amount float32) {
    s.loan += amount // (again and again and...) 
} 

//Employee实现SpendSalary方法 
func (e *Employee) SpendSalary(amount float32) {    
	e.money -= amount // More vodka please!!! Get me through the day! 
} 

// 定义interface 
type Men interface {    
	SayHi()    
	Sing(lyrics string)    
	Guzzle(beerStein string) 
} 

type YoungChap interface {    
	SayHi()    
	Sing(song string)    
	BorrowMoney(amount float32) 
} 

type ElderlyGent interface {    
	SayHi()    
	Sing(song string)    
	SpendSalary(amount float32) 
} 


func main() {
	xixi := Student{Human{"xixi",22,"123456789"},"台北一中",99.9}
	sisi := Employee{Human{"sisi",22,"1289"},"另一互联",10000}
	dia := Student{Human{"dia",22,"1789"},"台北一中",99.9}
	var m Men
	m = &xixi
	m.SayHi()
	m.Sing("可不可以")
	m.Guzzle("青岛纯生")

	var y YoungChap
	y = &dia
	y.SayHi()
	y.Sing("你好")
	y.BorrowMoney(100.0)

	var e ElderlyGent
	e = &sisi
	e.SayHi()
	e.Sing("爱你一万年")
	e.SpendSalary(900)
}