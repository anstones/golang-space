package main

import (
	"fmt"
	"math"
)


func main() {

	var x int = 100
   	var y int= 200

   	swap_(x,y)
   	fmt.Println(x)
   	fmt.Println(y)

   	swap__(&x, &y)
   	fmt.Println(x)
   	fmt.Println(y)
}

func max(num1, num2 int) int{
   /* 声明局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result 
}

func swap(x, y string) (string, string) {
   return y, x
}

// 值传递
func swap_(x, y int) int {
   var temp int
   temp = x /* 保存 x 的值 */
   x = y    /* 将 y 值赋给 x */
   y = temp /* 将 temp 值赋给 y*/
   return temp;
}

// 引用传递
func swap__(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}


// 函数定义后可作为值来使用
func func_value(){
   /* 声明函数变量 */
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* 使用函数 */
   fmt.Println(getSquareRoot(9))

}

