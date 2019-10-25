package main

import (
	"fmt"
)

func main() {
	 /* 数组长度为 5 */
   var  balance = []int {1000, 2, 3, 17, 50}

   /* 数组作为参数传递给函数 */
   avg := getAverage( balance, 5 )

   /* 输出返回的平均值 */
   fmt.Printf( "平均值为: %f ", avg )

   var array = []int{1, 2, 3, 4, 5}
    /* 未定义长度的数组只能传给不限制数组长度的函数 */
    setArray(array)
    /* 定义了长度的数组只能传给限制了相同数组长度的函数 */
    var array2 = [5]int{1, 2, 3, 4, 5}
    setArray2(array2)
	
}

func getAverage(arr []int, size int) float32 {
   var i,sum int
   var avg float32  

   for i = 0; i < size;i++ {
      sum += arr[i]
   }

   avg = float32(sum) / float32(size)

   return avg
}

func setArray(params []int) {
    fmt.Println("params array length of setArray is : ", len(params))
}

func setArray2(params [5]int) {
    fmt.Println("params array length of setArray2 is : ", len(params))
}