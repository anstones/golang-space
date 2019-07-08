package controllers

import (
	"encoding/json"
	. "firstapp/common/logger"
	"firstapp/models"
	"firstapp/runconfig"
	"firstapp/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


//  /user 请求的回复数据
type UserResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   models.User `json:"data"`
}

type Avatar struct {
	Url string `json:"avatar_url"`
}

// /user/avatar 请求的回复数据
type AvatarResp struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   Avatar `json:"data"`
}

type Name struct {
	Name string `json:"name"`
}

// /user/name 请求的回复数据
type NameResp struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   Name   `json:"data"`
}

type AuthInfo struct {
	RealName string `json:"real_name"`
	Id_card  string `json:"id_card"`
}

// /user/auth [POST] 请求的回复数据
type AuthResp struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}

type Houses struct {
	Houses []interface{} `json:"houses"`
}

// /user/Houses [GET] 请求的回复数据
type UserHousesResp struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   Houses `json:"data"`
}

type OrdersResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}


type UserHouseController struct {
	beego.Controller
}

func (this *UserHouseController) RetDate(rep interface{}){
	this.Data["json"] =rep
	this.ServeJSON()
}

func (this *UserHouseController) Get(){
	var rep = UserResp{Errno:utils.RECODE_OK,Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetDate(&rep)

	user_id := this.GetSession("user_id")
	if user_id == nil{
		rep.Errno = utils.RECODE_SESSIONERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}
	user := models.User{Id:user_id.(int)}
	o := orm.NewOrm()
	err := o.Read(user)
	if err != orm.ErrNoRows{
		Log.Debug("%s",err)
		rep.Errno = utils.RECODE_NODATA
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}else if err == orm.ErrMissPK{
		Log.Debug("%s",err)
		rep.Errno = utils.RECODE_NODATA
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}

	Log.Info("user_id:%s",user.Id, "user.name:%s",user.Name)
	user.Avatar_url = utils.AddDomain2Url(user.Avatar_url)

	rep.Data = user
	return
}

func (this *UserHouseController) Avatar(){
	var rep = AvatarResp{Errno:utils.RECODE_OK,Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetDate(&rep)

	file, header, err := this.GetFile("activer")
	if err != nil{
		Log.Debug("%s",err)
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}
	defer file.Close()

    // 保存文件到本地
	this.SaveToFile("avatar", "static/upload/" + header.Filename)

	activer_url := "htttp://" + runconfig.Config.Main.Httpaddr +":" + runconfig.Config.Main.Httpport + "static/upload/" + header.Filename
	user_id := this.GetSession("user_id")
	o := orm.NewOrm()
	user := models.User{Id:user_id.(int), Avatar_url:activer_url}

	//将用户上传的头像 写入数据库
	if _, err := o.Update(&user, "avatar_url"); err != nil {
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}

	rep.Data.Url = activer_url
	return
}

func (this *UserHouseController) UpdateName(){
	var rep = NameResp{Errno:utils.RECODE_OK,Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetDate(&rep)

	user_id := this.GetSession("user_id")

	var req Name

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req);err != nil{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}
	if req.Name == ""{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("name is Empty!")
		return
	}

	o := orm.NewOrm()
	user := models.User{Id:user_id.(int),Name:req.Name}

	if _,err := o.Update(&user, "name"); err != nil{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}

	//更新session
	this.SetSession("user_id", user_id)
	this.SetSession("name",req.Name)

	rep.Data = req
}

func (this *UserHouseController) AuthPost(){
	var rep = AuthResp{Errno:utils.RECODE_OK,Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetDate(&rep)

	user_id := this.GetSession("user_id")
	var req AuthInfo

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req);err != nil{
		Log.Debug("%s",err)
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}
	if req.Id_card == "" || req.RealName == ""{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("name or card is Empty!")
		return
	}
	o := orm.NewOrm()
	user := models.User{Id:user_id.(int),Real_name:req.RealName, Id_card:req.Id_card}

	if _, err := o.Update(&user,"real_name", "id_card");err != nil{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s",err)
		return
	}

	this.SetSession("user_id",user_id)

}

func (this *UserHouseController) GetHouses(){
	var rep = UserHousesResp{Errno:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetDate(&rep)

	user_id := this.GetSession("user_id")

	o := orm.NewOrm()
	qs := o.QueryTable("house")

	house_list := []models.House{}

	qs.Filter("user_id",user_id).RelatedSel().All(house_list)
	houses_rep := Houses{}
	for _,houseinfo := range house_list{
		fmt.Printf("house.user", houseinfo.User)
		fmt.Printf("house.area", houseinfo.Area)
		houses_rep.Houses = append(houses_rep.Houses, houseinfo.To_house_info())
	}

	fmt.Printf("house_rep: %+v\n",houses_rep)
	rep.Data = houses_rep
	return
}

func (this *UserHouseController) GetOrders() {
	var rep= OrdersResp{Errno: utils.RECODE_OK, Errmsg: utils.RecodeText(utils.RECODE_OK)}
	defer this.RetDate(&rep)

	user_id := this.GetSession("user_id")
	var role string
	this.Ctx.Input.Bind(&role, "role")

	if role == "" {
		rep.Errno = utils.RECODE_ROLEERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}
	o := orm.NewOrm()
	orders := []models.OrderHouse{}
	order_list := []interface{}{}

	if "landlord" == role {
		// 角色为房东
		landLordHouses := []models.House{}
		o.QueryTable("house").Filter("user__id", user_id).All(&landLordHouses)
		HouseIds := []int{}

		for _, house := range landLordHouses {
			HouseIds = append(HouseIds, house.Id)
		}
		o.QueryTable("order_house").Filter("house__id__in", HouseIds).OrderBy("-ctime").All(&orders)

	} else { // 角色为租客
		o.QueryTable("order_house").Filter("user__id", user_id).OrderBy("-ctime").All(&orders)
	}

	for _, order := range orders {
		o.LoadRelated(&order, "User")
		o.LoadRelated(&order, "House")
		order_list = append(order_list, order.To_order_info())
	}

	data := map[string]interface{}{}
	data["order"] = order_list

	rep.Data = data
	return
}

