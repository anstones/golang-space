package controllers

import (
	"encoding/json"
	. "firstapp/common/logger"
	"firstapp/models"
	"firstapp/runconfig"
	"firstapp/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type OrdersController struct {
	beego.Controller
}

type OrderRequest struct {
	House_id  	string 	`json:"house_id"`
	Start_data  string  `json:"start_data"`
	End_data    string  `json:"end_data"`
}

type OrderResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data,omitempty"`
}

func (this *OrdersController) RetData(rep interface{}) {
	this.Data["json"] = rep
	this.ServeJSON()
}

// /orders [POST]
func (this *OrdersController) Post()  {
	var rep = OrderResp{Errno:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)

	user_id := this.GetSession("user_id")

	var req OrderRequest
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req);err != nil{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}
	fmt.Printf("%+v\n",req)

	if req.House_id == "" || req.Start_data == "" || req.End_data == ""{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = "请求参数为空"
		return
	}
	//格式化时间
	start_date_time,_ := time.Parse("2006-01-02 15:04:05", req.Start_data+" 00:00:00")
	end_date_time,_ := time.Parse("2006-01-02 15:04:05", req.End_data+" 00:00:00")

	if end_date_time.Before(start_date_time){
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = "结束时间不能在开始时间之前"
		return
	}

	fmt.Printf("start_date_time = %+v, end_date_time = %+v\n", start_date_time, end_date_time)
	//计算入住天数
	days := end_date_time.Sub(start_date_time).Hours()/24 +1
	fmt.Printf("day is %f\n",days)

	house_id,_ := strconv.Atoi(req.House_id)
	house := models.House{Id:house_id}
	o := orm.NewOrm()
	if err := o.Read(&house);err != nil{
		rep.Errno = utils.RECODE_NODATA
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}
	o.LoadRelated(&house,"User")
	//房东不能够预定自己的房子
	if user_id == house.User.Id{
		rep.Errno = utils.RECODE_ROLEERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}
	// TODO 确保用户选择的房屋未被预定,日期没有冲突
	amount := days*float64(house.Price)
	order := models.OrderHouse{}
	order.House = &house
	user := models.User{Id:user_id.(int)}
	order.User = &user
	order.Begin_date = start_date_time
	order.End_date = end_date_time
	order.Days = int(days)
	order.House_price = house.Price
	order.Amount = int(amount)
	order.Status = models.ORDER_STATUS_WAIT_ACCEPT

	fmt.Printf("%+v\n", order)

	if _,err := o.Insert(&order);err != nil{
		rep.Errno = utils.RECODE_DBERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}
	// 更新session
	this.SetSession("user_id", user_id)
	data := map[string]interface{}{}
	data["order_id"] = order.Id
	rep.Data = data
	return
}

// /orders/:id/status [PUT]
func (this *OrdersController) OrderStatus()  {
	var rep = OrderResp{Errno:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)

	order_id := this.Ctx.Input.Param(":id")
	user_id := this.GetSession("user_id")
	var req map[string] interface{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err !=nil {
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}

	action := req["action"]
	if action != "accept" || action != "reject"{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}
	o := orm.NewOrm()
	order := models.OrderHouse{}
	if err := o.QueryTable("order_house").Filter("id",order_id).Filter("status", models.ORDER_STATUS_WAIT_ACCEPT).One(&order);err != nil{
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}

	if _,err := o.LoadRelated(&order, "House");err != nil{
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}
	house := order.House
	if house.User.Id != user_id{
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = "订单用户不匹配,操作无效"
		return
	}
	if action == "accept"{
		//如果是接受订单,将订单状态变成待评价状态
		order.Comment = models.ORDER_STATUS_WAIT_COMMENT
		Log.Info("action = accpet!")
	}else if action == "reject"{
		order.Status = models.ORDER_STATUS_REJECTED
		resone := req["reason"]
		order.Comment = resone.(string)
		Log.Info("action = accpet! the reason is: %s",resone)
	}
	// 更新数据库
	if _, err := o.Update(&order);err != nil{
		rep.Errno = utils.RECODE_DBERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}
	return
}

// /orders/:id/comment [PUT]
func (this *OrdersController) OrderComment()  {
	var rep = OrderResp{Errno:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)

	order_id := this.Ctx.Input.Param(":id")
	user_id := this.GetSession("user_id")
	// 获取参数
	var req map[string] interface{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req);err != nil{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s", err)
		return
	}
	comment := req["comment"].(string)
	if comment == ""{
		rep.Errno = utils.RECODE_PARAMERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}
	// 查询该订单，
	order := models.OrderHouse{}
	o := orm.NewOrm()
	if err := o.QueryTable("order_house").Filter("id",order_id).Filter("status",models.ORDER_STATUS_WAIT_COMMENT).One(&order);err != nil{
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}
    // 关联查询order订单所关联的user信息
	if _,err := o.LoadRelated(&order, "User");err != nil{
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}
	// 确保评价的人和该订单是同一个人
	if user_id != order.User.Id{
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = "该订单并不属于本人"
		return
	}
	if _,err := o.LoadRelated(&order, "House");err != nil{
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}

	house := order.House
	// 修改订单状态为： 已完成
	order.Status = models.ORDER_STATUS_COMPLETE
	order.Comment = comment

    // 更新该房屋的订单数量
	house.Order_count ++

	//更新house表和order表
	if _,err := o.Update(&house,"Order_count", ); err != nil{
		Log.Error("update house order_count error, err = ", err)
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}

	if _,err := o.Update(&order, "Status", "Comment");err != nil{
		Log.Error("updat order error, err = ", err)
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}

	// 删除redis缓存
	redis_config_map := map[string]string{
		"key": "ihome",
		"conn": runconfig.Config.Redis.Redisaddr + ":" + runconfig.Config.Redis.Redisport,
		"dbNum": runconfig.Config.Redis.Redisdbnum,
	}
	redis_config,_ := json.Marshal(redis_config_map)
	Log.Info("====%s",redis_config)

	cache_conn,err := cache.NewCache("redis", string(redis_config))
	if err != nil{
		Log.Debug("connect cache error %s", err)
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return
	}
	house_info_key := "house_info_" + strconv.Itoa(house.Id)
	if err := cache_conn.Delete(house_info_key);err != nil{
		Log.Debug("delete ", house_info_key, "error , err = ", err)
	}
	return
}