package config

import (
	"github.com/astaxie/beego"
)

const (
	DbAddr string = "root:root@/server?charset=utf8&parseTime=True&loc=Local"
	IMAGE_FILE_PATH = "../wts_web/static/upload/"
)

func InitConfig() {
	//app config
	beego.BConfig.AppName = "server"
	beego.BConfig.RunMode = "dev"
	beego.BConfig.RouterCaseSensitive = true
	beego.BConfig.RecoverPanic = false
	beego.BConfig.MaxMemory = 1 << 26
	beego.BConfig.EnableErrorsShow = true
	beego.BConfig.CopyRequestBody = true

	//listenning config
	beego.BConfig.Listen.ServerTimeOut = 120
	beego.BConfig.Listen.HTTPPort = 8080
	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"

	//web config
	beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.BConfig.WebConfig.XSRFExpire = 24 * 60 * 60

	//session config
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionName = "session_id"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true

	//log config
	beego.BConfig.Log.AccessLogs = true
	beego.BConfig.Log.FileLineNum = true

}
