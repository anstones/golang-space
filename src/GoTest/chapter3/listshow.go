package chapter3

import (
	"container/list"
	"fmt"
)

func main()  {

	l := list.New()
	l.PushBack("canon")
	l.PushFront(67)

	element := l.PushBack("fist")

	l.InsertAfter("high", element)

	l.InsertBefore("noon",element)

	l.Remove(element)

	for i := l.Front(); i != nil;i = i.Next(){  //Front 获取列表头元素， Next 获取下一个元素
		fmt.Println(i.Value)
	}
}
