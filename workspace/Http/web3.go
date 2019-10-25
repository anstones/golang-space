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
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
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
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
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

