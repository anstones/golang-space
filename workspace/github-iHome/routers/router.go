package routers

import (
	"firstapp/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

//func init() {
//  beego.Router("/", &controllers.MainController{})
//	beego.Router("/user", &controllers.UserController{})
//	beego.Router("/test", &controllers.UserController{},"get:Test")
//	beego.Router("/post", &controllers.UserController{},"post:TestPost")
//	beego.Router("/image", &controllers.UserController{},"post:PostImage")
//}

func ignoreStaticPath()  {
	// 透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context)  {
	orpath := ctx.Request.URL.Path
	if strings.Index(orpath, "api") >0{
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/" + ctx.Request.URL.Path)
}

func init()  {
	ignoreStaticPath()
	ns :=
		beego.NewNamespace("/api",
			//beego.NSCond(func(ctx *context.Context) bool {
			//	if ctx.Input.Domain() == "192.168.6.66"{
			//		fmt.Println("=========================")
			//		return true
			//	}
			//	loggers.Info("now domaain is:", ctx.Input.Domain())
			//	return false
			//}),
			beego.NSNamespace("/v1.0",
				//注册
				beego.NSRouter("/users", &controllers.RegController{}),
				//登陆
				beego.NSRouter("/sessions", &controllers.LoginController{}),
				//请求地理区域信息
				beego.NSRouter("/areas", &controllers.AreaController{}),
				//验证用户是否已经注册
				beego.NSRouter("/session", &controllers.SessionController{}),


				//请求当前用户信息
				beego.NSRouter("/user", &controllers.UserHouseController{}),
				//用户上传用户头像
				beego.NSRouter("/user/avatar", &controllers.UserHouseController{}, "post:Avatar"),
				//用户更新用户名
				beego.NSRouter("/user/name", &controllers.UserHouseController{}, "put:UpdateName"),
				//请求用户身份认证信息, 上传用户身份信息
				beego.NSRouter("/user/auth", &controllers.UserHouseController{}, "get:Get;post:AuthPost"),
				//请求当前用户的所有发布的房源信息列表
				beego.NSRouter("/user/houses", &controllers.UserHouseController{}, "get:GetHouses"),
				//请求当前用户提交的订单列表信息
				beego.NSRouter("/user/orders", &controllers.UserHouseController{}, "get:GetOrders"),


				//用户发布房源信息
				beego.NSRouter("/houses", &controllers.HousesController{}, "post:Post;get:Get"),
				//用户上传房源图片
				beego.NSRouter("/houses/:id/images", &controllers.HousesController{}, "post:UploadHouseImage"),
				//用户请求房源详细信息
				beego.NSRouter("/houses/:id", &controllers.HousesController{}, "get:GetOneHouseInfo"),
				//用户请求房源首页列表信息
				beego.NSRouter("/houses/index", &controllers.HousesController{}, "get:IndexHouses"),


				//用户下单请求
				beego.NSRouter("/orders", &controllers.OrdersController{}, "post:Post"),
				//房东用户接受/拒绝 订单请求
				beego.NSRouter("/orders/:id/status", &controllers.OrdersController{}, "put:OrderStatus"),
				//用户发送订单评价信息
				beego.NSRouter("/orders/:id/comment", &controllers.OrdersController{}, "put:OrderComment"),
			),
	)
	beego.AddNamespace(ns)
}
