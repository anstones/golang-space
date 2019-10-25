package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName (w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()  // 解析参数， 默认是不解析
	fmt.Println("========",r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key", k)
		fmt.Println("value", strings.Join(v,""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main()  {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		log.Fatal("ListenAndServer:", err)
	}
}