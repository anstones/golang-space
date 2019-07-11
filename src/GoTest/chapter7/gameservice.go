package main

//  一个类型可以实现多个接口
//  一个接口的方法，不一定需要一个类型完全实现，可以通过在类型中嵌入其他类型来是实现

import "fmt"

type Service interface {
	Start()
	Log(string)
}

type Logger struct {
}

func (g Logger) Log(l string) {
	fmt.Println(l)
}

type GameService struct {
	Logger //嵌入日志器， 日志器实现了 loggers 方法
}

func (g GameService) Start() {

}

func main() {
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")
}
