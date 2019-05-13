package main

import "fmt"

// 空接口使用
// 字典的 key 是字符串，但是希望 value 可以容纳任意类型的对象，
// 类似于 Java 语言的 Map 类型，这时候就可以使用空接口类型 interface{}
func main() {
 // 连续两个大括号，是不是看起来很别扭
    var user = map[string]interface{}{
        "age": 30,
        "address": "Beijing Tongzhou",
        "married": true,
    }
    fmt.Println(user)
    // 类型转换语法来了
 var age = user["age"].(int)
    var address = user["address"].(string)
    var married = user["married"].(bool)
    fmt.Println(age, address, married)
}
