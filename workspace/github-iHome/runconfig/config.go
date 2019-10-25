package runconfig

import (
	"github.com/astaxie/beego"
)


type logParam struct {
	FilePath     string
	Level        int
}

type mysqlParam struct {
	Mysqladdr    string
	Mysqlport    string
	Mysqluser    string
	Mysqlpswd    string
	Mysqldb  string
}

type redisParam struct {
	Redisaddr    string
	Redisport    string
	Redisdbnum   string
}

type mainServer struct {
	Httpaddr   string
	Httpport   string
}

type configParam struct {
	Log     	logParam
	Mysql   	mysqlParam
	Redis   	redisParam
	Main        mainServer
}

var Config = &configParam{}

func (c *configParam) Parse() error{

	c.Log.FilePath = beego.AppConfig.DefaultString("log::filepath", "./log/data_analysis_service.log")
	c.Log.Level = beego.AppConfig.DefaultInt("log::loglevel", 7)

	c.Mysql.Mysqladdr = beego.AppConfig.DefaultString("mysql::mysqladdr", "192.168.133.128")
	c.Mysql.Mysqlport= beego.AppConfig.DefaultString("mysql::mysqlport", "3306")
	c.Mysql.Mysqluser = beego.AppConfig.DefaultString("mysql::mysqluser", "root")
	c.Mysql.Mysqlpswd = beego.AppConfig.DefaultString("mysql::mysqlpswd", "mysql")
	c.Mysql.Mysqldb = beego.AppConfig.DefaultString("mysql::mysqldbname","ihome")

	c.Redis.Redisaddr = beego.AppConfig.DefaultString("mysql::redisaddr", "192.168.133.128")
	c.Redis.Redisport = beego.AppConfig.DefaultString("mysql::redisport", "6379")
	c.Redis.Redisdbnum = beego.AppConfig.DefaultString("mysql::redisdbnum", "1")

	c.Main.Httpaddr = beego.AppConfig.DefaultString("mainserver::httpaddr", "192.168.6.66")
	c.Main.Httpport = beego.AppConfig.DefaultString("mainserver::httpport", "8080")
	return nil
}