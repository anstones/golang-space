package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	. "firstapp/common/logger"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

type User struct {
	Name string
	Pwd  string
	Age  int
	Sex  string
}

type  Mystruct struct {
	Name string
	Age  int
}


func (c *UserController) Test(){
	username := c.GetString("username")
	fmt.Println(username)
	//c.Ctx.WriteString("hello world")
	mystruct := &Mystruct{username, 34}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}

func (this *UserController) Get(){
	this.Ctx.WriteString("hello")  //不调用beego 的Render 函数 ，直接返回
}

func (this *UserController) TestPost(){
	user := User{}
	if error:=this.ParseForm(&user);error!=nil {
		this.Ctx.WriteString("出错了！")
		Log.Debug("beego error", error)
	}else{
		this.Ctx.WriteString("Name="+user.Name+"\nPwd="+user.Pwd+"\nSex="+user.Sex+"\nAge="+strconv.Itoa(user.Age))
	}
}

func (this *UserController) PostImage(){
	f, h, err := this.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	this.SaveToFile("uploadname", "static/upload/" + h.Filename)
	this.Ctx.WriteString("upload secssed")
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	Log.Info("My frist beego app start")
}



