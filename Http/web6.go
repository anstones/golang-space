// 防止form表单重复提交

package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"html/template"
	"strconv"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("C:\\Users\\Administrator\\Desktop\\mine\\Project\\Go_Space\\Http\\login.gtpl")
		t.Execute(w, token)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
		 }
}

func main()  {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9000", nil)
	if err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}