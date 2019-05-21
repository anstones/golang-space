package main

import(
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"strconv"
	"time"
)

func say(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.Method)
	for k, v := range r.Form{
		fmt.Println("kwy:", k)
		fmt.Println("val:", strings.Join(v,""))
	}
	fmt.Fprintf(w, r.Method)
}

func login(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	// 验证表单： 必填字段
	if len(r.Form["username"][0]) == 0{
		//为空时的处理
	}

	//验证表单：数字
	getint,err := strconv.Atoi(r.Form.Get("age"))
	if err != nil{
		//数字转化出错，那么可能不是数字
	}
	if getint > 100{
		//太大了
	}
	// 正则判断数字
	if m,_ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m{
		return false
	}

	//验证表单：中文
	if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("realname")); !m {
		return false
	}

	//验证表单：英文
	if m, err := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
		fmt.Println("error:", err)
		return false
	}

	//验证表单：邮箱
	if m, err:= regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m{
		fmt.Println("error:", err)
	}else{
		fmt.Println("yes") }

	//验证表单：手机号码
	if m, err := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
		fmt.Println("error:", err)
	}else{
		fmt.Println("yes")
	}

	//验证表单：下拉菜单
	slice := []string{"apple","pear","banane"}
	for _, v := range slice {
		if v == r.Form.Get("fruit") {
			return true
		}
	}
	return false

	//验证表单： 单选
	slice := []int{1,2}
	for _, v := range slice {
		if v == r.Form.Get("gender") {
			return true
		}
	}
	return false

	//验证表单：复选框
	slice:=[]string{"football","basketball","tennis"}
	a := Slice_diff(r.Form["interest"], slice)
	if a == nil{
		return true
	}
	return false

	//验证表单： 时间
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

	//验证表单：身份证
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		return false
	}
}

func main()  {
	http.HandleFunc("/", say)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe error: ", err)
	}
}