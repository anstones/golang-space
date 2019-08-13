package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func (this *Person) String() string {
	buffer := bytes.NewBufferString("这是 ")
	buffer.WriteString(this.Name + ",")
	if this.Sex == 0 {
		buffer.WriteString("他")
	} else {
		buffer.WriteString("她")
	}
	buffer.WriteString("今年 ")
	buffer.WriteString(strconv.Itoa(this.Age))
	buffer.WriteString(" 岁")
	return buffer.String()
}

func (this *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(this.String()))
		f.Write([]byte("。Person有三个字段"))
	} else {
		f.Write([]byte(fmt.Sprintln(this.String())))
	}
}

func main() {
	P := new(Person)
	P.Age = 28
	P.Sex = 0
	P.Name = "zuolan"
	//P := &Person{"zuolan",28,0}
	//fmt.Println(P)
	fmt.Printf("%L", P)
}
