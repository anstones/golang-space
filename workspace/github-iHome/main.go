package main

import (
	. "firstapp/common/logger"
	"firstapp/initializer"
	_ "firstapp/models"
	_ "firstapp/routers"
	_ "firstapp/runconfig"
	"github.com/astaxie/beego"
)

func main() {

	LogTemp("DAS begin to start")
	if err := initializer.Initlize(); err != nil {
		return
	}

	beego.SetStaticPath("/static", "static") //添加静态文件目录
	beego.Run()
}



