package main

import "fmt"

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type Bird struct {
}

func (b *Bird) Fly() {
	fmt.Println("bird : fly")
}

func (p *Bird) Walk() {
	fmt.Println("bird : walk")
}

type Pig struct {
}

func (p *Pig) Walk() {
	fmt.Println("pig : walk")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(Bird),
		"pig":  new(Pig),
	}

	for name, obj := range animals {
		f, isFlyer := obj.(Flyer) // 类型断言 t, ok := i.(T)  i：接口变量；T: 转换的目标类型； t: 转换后的变量
		w, isWalker := obj.(Walker)

		fmt.Printf("name: %s isFlyer :%v isWalker: %v\n", name, isFlyer, isWalker)

		if isFlyer {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}
	}
}
