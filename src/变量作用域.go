package main

import "fmt"

func main(){
  var a int = 0
  fmt.Println("for start")
  for a:=0; a < 10; a++ {
    fmt.Println(a)
  }
  fmt.Println("for end")
  fmt.Println(a)
}


func main(){ 
	var a int = 0 
	fmt.Println("for start") 
	for a = 0; a < 10; a++ { 
		fmt.Println(a) 
	} 
	fmt.Println("for end") 
	fmt.Println(a) 
}