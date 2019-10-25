package controllers

import (
	"encoding/json"
	. "firstapp/common/logger"
	"firstapp/models"
	"firstapp/runconfig"
	"firstapp/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"time"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(rep interface{}) {
	this.Data["json"] = rep
	this.ServeJSON()
}

type AreaResp struct {
	Error  string `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func (this *AreaController) Get()  {
	var rep = AreaResp{Error:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)

	// 1. 从redis读取地域信息缓存，如果有就直接返回
	areas_info_key := "area_info"
	redis_config_map := map[string]string{
		"key" :"ihome",
		"conn": runconfig.Config.Redis.Redisaddr +":"+ runconfig.Config.Redis.Redisport,
		"dbNum":runconfig.Config.Redis.Redisdbnum,
	}
	redis_config,_ := json.Marshal(redis_config_map)
	Log.Info("====%s",redis_config)
	cache_conn,err := cache.NewCache("redis", string(redis_config))
	if err != nil{
		Log.Debug("connect cache error %s", err)
		rep.Error = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return
	}

	areas_info_value := cache_conn.Get(areas_info_key)
	if areas_info_value != nil{
		var areas_info interface{}
		_ = json.Unmarshal(areas_info_value.([]byte), &areas_info)
		rep.Data = areas_info
		return
	}

	// 2. 如果redis缓存中没有，从mysql查询，然后存入redis
	o := orm.NewOrm()
	var areas []models.Area
	num,err := o.QueryTable("area").All(&areas)
	if err != nil {
		rep.Error = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return
	}
	if num == 0{
		rep.Error = utils.RECODE_NODATA
		rep.Errmsg = utils.RecodeText(utils.RECODE_NODATA)
		return
	}

	// 3. 将信息返回
	area_list := []models.Area{}
	for _,area := range areas{
		Log.Info("%+v\n", area)
		area_list = append(area_list,area)
	}
	areas_info_value, _ = json.Marshal(area_list)
	cache_conn.Put(areas_info_key,areas_info_value,3600*time.Second)
	rep.Data = area_list
	return
}
