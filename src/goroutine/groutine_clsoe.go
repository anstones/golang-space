package main

import (
	"fmt"
)

func SendDate(ch chan string)  {
	ch <- "纽约"
	ch <- "华盛顿"
	ch <- "伦敦"
	ch <- "北京"
	ch <- "上海"
	close(ch)
}

func GetDtae(ch chan string)  {
	for {
		if v,ok := <- ch;ok{
			fmt.Println(v)
		}else {
			break
		}
	}
}


func main()  {
	ch := make(chan string)
	go SendDate(ch)
	GetDtae(ch)
}
