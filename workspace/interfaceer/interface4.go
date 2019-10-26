package main

import (
	"fmt"
)

// 接口中是方法的定义
type Pet interface {
	SetName(name string)
	Name() string
}

type Dog struct {
	name string
}

// *Dog 的方法集合中有 Pet 接口的全部方法
// *Dog 就是 Pet 接口类型的实现类型，这种无侵入式的接口实现方式叫”鸭子类型“
func (dog *Dog) SetName(name string) {
	dog.name = name
}

// Dog 的方法集合中仅有 Pet 接口中的 Name 方法，所以没有实现 Pet 接口
func (dog Dog) Name() string {
	return dog.name
}

func main(){
	dog := Dog {"little dog"}
	_, ok := interface{}(dog).(Pet)   // 不能转换
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)  // 可以转换
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	// pet 的静态类型为 Pet，动态类型为 *Dog
	var pet Pet = &dog   // 赋值成功，因为 *Dog 是 Pet 的实现
	fmt.Printf("name: %q.\n", pet.Name())
}
