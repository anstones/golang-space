package main

import "fmt"

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main() {
   var numbers []int
   printSlice(numbers)

   /* 允许追加空切片 */
   numbers = append(numbers, 0)
   printSlice(numbers)

   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   printSlice(numbers)

   /* 同时添加多个元素 */
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)

   /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)

   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)

   veggies := []string{"potatoes", "tomatoes", "brinjal"}
   fruits := []string{"oranges", "apples"}

   fmt.Println("len",len(veggies), "cap", cap(veggies))
   fmt.Println("len",len(fruits), "cap", cap(fruits))
   food := append(veggies, fruits...)
   fmt.Println("food:",food)

   fmt.Println("len",len(food), "cap", cap(food))


   darr := []int{57, 89, 90, 82, 100, 78, 67, 69, 59}
   darr = append(darr, 100)

   fmt.Println("len",len(darr), "cap", cap(darr))
   dslice := darr[2:5]
   fmt.Println("array before", darr)
   for i := range dslice {
      dslice[i]++
   }
   fmt.Println("array after", darr)
}

