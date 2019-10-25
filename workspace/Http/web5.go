// 防止跨域请求

//  func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
//  func HTMLEscapeString(s string) string //转义s之后返回结果字符串
//  func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)


func sayhelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()   //解析url传递的参数，对于post则解析响应包的主题（request body)
	// 如果没有parseForm方法。下面无法获取表单的数据
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	for k, v := range r.Form{
		fmt.Println("key:", k)
		fmt.Println("val;",strings.Join(v,""))
	}
	fmt.Fprintf(w, "hello world")
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method)
	if r.Method == "GET"{
		t, _ := template.ParseFiles("C:\\Users\\Administrator\\Desktop\\mine\\Project\\Go_Space\\Http\\login.gtpl")
		t.Execute(w, nil)
	}else {
		r.ParseForm()
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

func main()  {
	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9000", nil)
	if err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}


