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

type HouseInfo struct {
	Area_id    string   `json:"area_id"`    //归属地的区域编号
	Title      string   `json:"title"`      //房屋标题
	Price      string   `json:"price"`      //单价,单位:分
	Address    string   `json:"address"`    //地址
	Room_count string   `json:"room_count"` //房间数目
	Acreage    string   `json:"acreage"`    //房屋总面积
	Unit       string   `json:"unit"`       //房屋单元,如 几室几厅
	Capacity   string   `json:"capacity"`   //房屋容纳的总人数
	Beds       string   `json:"beds"`       //房屋床铺的配置
	Deposit    string   `json:"deposit"`    //押金
	Min_days   string   `json:"min_days"`   //最好入住的天数
	Max_days   string   `json:"max_days"`   //最多入住的天数 0表示不限制
	Facilities []string `json:"facility"`   //房屋设施
}


type House_id struct {
	House_id int64 `json:"house_id"`
}

type HousesResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

type Image_url struct {
	Url    string      `json:"url"`
}

type HousesController struct {
	beego.Controller
}

func (this *HousesController) RetData(rep interface{}) {
	this.Data["json"] = rep
	this.ServeJSON()
}

type HouseOneResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

type HouseIndexResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

type HouseImageResp struct {
	Errno  string    `json:"errno"`
	Errmsg string    `json:"errmsg"`
	Data   Image_url `json:"data"`
}


// /houese?aid=1&sd=2017-11-09&ed=2017-11-11&sk=new&p=1 [GET]
func (this *HousesController) Get() {
	rep := HouseOneResp{Errno: utils.RECODE_OK, Errmsg: utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)

	var aid int
	this.Ctx.Input.Bind(&aid, "aid")
	var sd string
	this.Ctx.Input.Bind(&sd, "sd")
	var ed string
	this.Ctx.Input.Bind(&ed, "ed")
	var sk string
	this.Ctx.Input.Bind(&sk, "sk")
	var page int
	this.Ctx.Input.Bind(&page, "p")

	Log.Debug("%d,%s,%s,%s,%d",aid, sd, ed, sk, page)

	//把时间从str转换成字符串格式

	//校验开始时间一定要早于结束时间

	//判断page的合法性 一定是大于0的整数

	//尝试从redis中获取数据, 有则返回

	//如果没有 从mysql中查询
	houses := []models.House{}

	o := orm.NewOrm()

	qs := o.QueryTable("house")

	num, err := qs.Filter("area_id", aid).All(&houses)
	if err != nil {
		rep.Errno = utils.RECODE_PARAMERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		return
	}

	total_page := int(num)/models.HOUSE_LIST_PAGE_CAPACITY + 1
	house_page := 1

	house_list := []interface{}{}
	for _, house := range houses {
		o.LoadRelated(&house, "Area")
		o.LoadRelated(&house, "User")
		o.LoadRelated(&house, "Images")
		o.LoadRelated(&house, "Facilities")
		house_list = append(house_list, house.To_house_info())
	}

	data := map[string]interface{}{}
	data["houses"] = house_list
	data["total_page"] = total_page
	data["current_page"] = house_page

	rep.Data = data

	return
}

// /houese [POST]  表当增加房屋 Post
func (this *HousesController) Post()  {
	 var rep = HousesResp{Errno:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	 defer this.RetData(&rep)

	 var req = HouseInfo{}
	 if err := json.Unmarshal(this.Ctx.Input.RequestBody,&req); err != nil{

		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("%s", err)
		return
	 }
	 fmt.Printf("%v+\n", req)
	 house := models.House{}
	 house.Room_count, _ = strconv.Atoi(req.Room_count)
	 house.Title = req.Title
	 house.Acreage, _ = strconv.Atoi(req.Acreage)
	 house.Unit = req.Unit
	 house.Deposit, _ = strconv.Atoi(req.Deposit)
	 house.Deposit = house.Deposit * 100 //单位转换
	 house.Address = req.Address
	 house.Price, _ = strconv.Atoi(req.Price)
	 house.Price = house.Price * 100 //单位转换
	 house.Capacity, _ = strconv.Atoi(req.Capacity)
	 house.Beds = req.Beds
	 house.Min_days, _ = strconv.Atoi(req.Min_days)
	 house.Max_days, _ = strconv.Atoi(req.Max_days)
	 user := models.User{Id: this.GetSession("user_id").(int)}
	 area_id, _ := strconv.Atoi(req.Area_id)
	 area := models.Area{Id: area_id}
	 house.User = &user
	 house.Area = &area
	 //插入house 到数据库
	 o := orm.NewOrm()
	 house_id, err := o.Insert(&house)
	 if err != nil{
		 rep.Errno = utils.RECODE_DBERR
		 rep.Errmsg = utils.RecodeText(rep.Errno)
		 Log.Debug("%s", err)
		 return
	 }
	 Log.Info("house insert id =", house_id, " succ!")

	 facilities := []*models.Facility{}
	 for _, fid := range req.Facilities{
	 	id, _ := strconv.Atoi(fid)
	 	facility := &models.Facility{Id:id}
		 facilities = append(facilities, facility)
	 }

	 m2mhouse_facility := o.QueryM2M(&house,"Facilities")

	 num, err := m2mhouse_facility.Add(facilities)
	 if err != nil{
		 rep.Errno = utils.RECODE_DBERR
		 rep.Errmsg = utils.RecodeText(rep.Errno)
		 Log.Debug("%s", err)
		 return
	 }
	 Log.Info("house m2m facility insert num =", num, " succ!")

	 //将index缓存删除
	 redis_config_map := map[string]string{
	 	"key":   "ihome_go",
	 	"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
	 	"dbNum": utils.G_redis_dbnum,}
	 redis_config, _ := json.Marshal(redis_config_map)
	 cache_conn, err := cache.NewCache("redis", string(redis_config))
	 if err != nil {
	 	Log.Debug("connect cache error")}
	 house_page_key := "home_page_data"
	 cache_conn.Delete(house_page_key)
	 rep.Data = House_id{House_id: house_id}
	 return
}

// /houses/index [GET]
func (this *HousesController) IndexHouses()  {
	var rep = HouseIndexResp{Errno:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)

	data := []interface{}{}
	Log.Debug("Index Houses....")

	//1 从缓存服务器中请求 "home_page_data" 字段,如果有值就直接返回
	//先从缓存中获取房屋数据,将缓存数据返回前端即可
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
		rep.Errno = utils.RECODE_DATAERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return
	}
	house_page_key := "home_page_data"
	house_page_value := cache_conn.Get(house_page_key)

	if house_page_value != nil{
		Log.Info("======= get house page info  from CACHE!!! ========")
		_ = json.Unmarshal(house_page_value.([]byte), &data)
		rep.Data = data
		return
	}
	//2 如果缓存没有,需要从数据库中查询到房屋列表
	houses := []models.House{}
	o := orm.NewOrm()

	if _,err := o.QueryTable("house").Limit(models.HOME_PAGE_MAX_HOUSES).All(&houses); err != nil{
		for _, house := range houses{
			o.LoadRelated(&house, "Area")
			o.LoadRelated(&house, "User")
			o.LoadRelated(&house, "Images")
			o.LoadRelated(&house, "Facilities")
			data = append(data, house.To_house_info())
		}
	}
	//将data存入缓存数据
	house_page_value,_ = json.Marshal(data)
	cache_conn.Put(house_page_key, house_page_value, 3600*time.Second)

	rep.Data = data
	return
}

// /houses/:id/images [POST]
func (this *HousesController) UploadHouseImage()  {
	var rep = HouseImageResp{Errno:utils.RECODE_OK, Errmsg:utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)

	house_id := this.Ctx.Input.Param(":id")
	file, header, err := this.GetFile("house_image")
	if err != nil{
		rep.Errno = utils.RECODE_REQERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_REQERR)
		Log.Debug("%s", err)
		return
	}
	defer file.Close()

	//将文件保存到本地
	this.SaveToFile("house_image", "static/upload/" + header.Filename)

	house_image_url := "http://" + runconfig.Config.Main.Httpaddr +":" + runconfig.Config.Main.Httpport + "static/upload/" + header.Filename

	house := models.House{}
	house.Id,_ = strconv.Atoi(house_id)

	o := orm.NewOrm()
	if err := o.Read(&house); err != nil{
		rep.Errno = utils.RECODE_DBERR
		rep.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		Log.Debug("%s",err)
		return
	}
	//根据house_id 查询house_image 是否为空
	if house.Index_image_url == ""{
		//如果为空 那么就用当前image_url为house的主image_url
		house.Index_image_url = house_image_url
		Log.Info("house Index_image_url set %s", house_image_url)
	}

	house_image := models.HouseImage{House:&house, Url:house_image_url}
	//将house_image 和hosue相关联
	house.Images = append(house.Images,&house_image)

	//将house_image入库
	if _, err := o.Insert(&house_image); err != nil{
		rep.Errno = utils.RECODE_DBERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("insert house image error")
		return
	}

	if _, err := o.Update(&house);err != nil{
		rep.Errno = utils.RECODE_DBERR
		rep.Errmsg = utils.RecodeText(rep.Errno)
		Log.Debug("update house error")
		return
	}
	image := Image_url{Url:house_image_url}
	rep.Data = image
	return
}

// /houses/:id [GET]
func (this *HousesController) GetOneHouseInfo()  {
	rep := HouseOneResp{Errno: utils.RECODE_OK, Errmsg: utils.RecodeText(utils.RECODE_OK)}
	defer this.RetData(&rep)

	data := make(map[string] interface{})

	user_id := this.GetSession("user_id")
	if user_id == nil{
		Log.Debug("user not login")
		user_id = -1
	}
	house_id := this.Ctx.Input.Param(":id")

	redis_config_map := map[string]string{
		"key": "ihome",
		"conn": runconfig.Config.Redis.Redisaddr + ":" + runconfig.Config.Redis.Redisport,
		"dbNum" : runconfig.Config.Redis.Redisdbnum,
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

	house_info_key := fmt.Sprintf("house_info_%s", house_id)
	house_info_value := cache_conn.Get(house_info_key)

	if house_info_value != nil{
		Log.Info("======= get house page info  from CACHE!!! ========")
		data["user_id"] = user_id
		house_info := map[string]interface{}{}
		json.Unmarshal(house_info_value.([]byte), &house_info)
		data["house"] = house_info
		rep.Data = data
		return
	}
	o := orm.NewOrm()

	// 载入关系查询 -----
	house := models.House{}
	house.Id,_ = strconv.Atoi(house_id)
	o.Read(&house)
	o.LoadRelated(&house, "Area")
	o.LoadRelated(&house, "user")
	o.LoadRelated(&house, "Images")
	o.LoadRelated(&house, "Facilities")

	// 存入缓存
	house_info_value,_ = json.Marshal(house.To_one_house_desc())
	cache_conn.Put(house_info_key, house_info_value, 3600*time.Second)

	data["user_id"] = user_id
	data["house"] = house.To_one_house_desc()

	rep.Data = data
	return
}