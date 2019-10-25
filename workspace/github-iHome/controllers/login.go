package controllers

import (
	"encoding/json"
	. "firstapp/common/logger"
	"firstapp/models"
	"firstapp/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


// 统一回复数据格式
type LoginResp struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}

type LoginReq struct {
	Moblie   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) RetData(rep interface{})  {
	c.Data["json"] = rep
	c.ServeJSON()
}

func (this *LoginController) Post()  {
	var req LoginReq

	rep := LoginResp{Errno:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(rep)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil{
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return
	}
	Log.Info("Login RequestInfo:%+v\n", req)

	if req.Moblie == "" || req.Password == ""{
		rep.Errno = utils.RECODE_PARAMERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_PARAMERR)
		return
	}
	var user models.User
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("mobile",req.Moblie).One(&user)
	if err == orm.ErrNoRows {
		//没有该用户
		rep.Errno = utils.RECODE_LOGINERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_LOGINERR)
		return
	}

	if user.PasswordHash != req.Password{
		rep.Errno = utils.RECODE_PWDERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_PWDERR)
		return
	}

	// 存入session,状态保持
	this.SetSession("user_id", user.Id)
	if user.Name == ""{
		this.SetSession("name", user.Mobile)
	}else {
		this.SetSession("name",user.Name)
	}
	this.SetSession("mobile", user.Mobile)
	return

}