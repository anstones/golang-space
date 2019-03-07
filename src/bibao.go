
package main

import (
  "fmt"
)


func main() {
  getSequence_func := getSequence()

  fmt.Println(getSequence())
  fmt.Println(getSequence_func()) //打印出getSequence内存地址

  add_func := add(1,2)
  fmt.Println(add_func())

  add_add_func := add_add(1,2)
  fmt.Println(add_add_func(0,1))
}


// 闭包

func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i  
   }
}


// 闭包带参数
func add(x1, x2 int) func()(int,int)  {
    i := 0
    return func() (int,int){
        i++
        return i,x1+x2
    }
}

// 闭包带参数1
func add_add(x1, x2 int) func(x3, x4 int)(int,int,int)  {
    i := 0
    return func(x3,x4 int) (int,int,int){ 
       i++
       return i,x1+x2,x3+x4
    }
}