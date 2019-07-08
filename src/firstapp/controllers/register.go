package controllers

import (
	"firstapp/models"
	"fmt"
	. "firstapp/common/logger"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"firstapp/utils"
)

// 统一回复数据格式
type RegResp struct {
	Error  string `json:"error"`
	Errmsg string `json:"errmsg"`
}

type RegReq struct {
	Moblie   string `json:"mobile"`
	Password string `json:"password"`
	Sms_code string `json:"sms_code"`
}

type RegController struct {
	beego.Controller
}

func (c *RegController) RetData(rep interface{})  {
	c.Data["json"] = rep
	c.ServeJSON()
}

func (c *RegController) Post()  {
	var req RegReq
	rep := RegResp{Error:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer c.RetData(&rep)

	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	fmt.Println("Reg RequestInfo:%+v\n", req)

	if req.Moblie == "" || req.Password == "" || req.Sms_code == ""{
		rep.Error = utils.RECODE_PARAMERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_PARAMERR)
		return
	}

	user := models.User{}
	user.Mobile = req.Moblie
	user.PasswordHash = req.Password

	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil{
		rep.Error = utils.RECODE_DBERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	Log.Info("reg insert id =", id, "succ!")

}