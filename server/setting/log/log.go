package logs

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	runmode := beego.AppConfig.String("runmode")
	logs.SetLogger("console")
	if runmode == "pro" {
		logs.SetLogger(logs.AdapterFile, `{"filename":"/web/log/server.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	}
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}
