package main

import "fmt"

func printMsgType(msg *struct{
	id int
	data string
})  {
	fmt.Printf("%T\n",msg)
	fmt.Printf("%d : %s", msg.id, msg.data)
}

func main()  {
	msg := struct {
		id int
		data string
	}{
		id: 1024,
		data: "hello",
	}

	printMsgType(&msg)
}