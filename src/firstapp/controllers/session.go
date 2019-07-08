package controllers
// 判断用户是否注册过

import (
	. "firstapp/common/logger"
	"firstapp/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type  SessionController struct {
	beego.Controller
}

type SessionResp struct {
	Error  string 				`json:"error"`
	Errmsg string 				`json:"errmsg"`
	Data   SessionUsername 		`json:"data"`
}

type SessionUsername struct {
	Name string 				`json:"name"`
}

func (this *SessionController) RetData(rep interface{}) {
	this.Data["json"] = rep
	this.ServeJSON()
}

func (this *SessionController) Get()  {
	var rep = SessionResp{Error:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)

	name := this.GetSession("name")
	Log.Info("name : %+v\n",name)
	if name == nil{
		rep.Error = utils.RECODE_SESSIONERR
		rep.Errmsg = utils.RecodeText(rep.Error)
		return
	}else{
		rep.Data = SessionUsername{Name:fmt.Sprintf("%s",name)}
		return
	}
}

func (this *SessionController) Delete()  {
	var rep = SessionResp{Error:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)
	this.DelSession("user_id")
	this.DelSession("name")
	return
}