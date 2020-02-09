package router

import (
	"github.com/astaxie/beego"
	account "server/account/controller"
)

func InitRouter() {
	beego.Router("api/LoginAPI", &account.LoginAPI{})
	beego.Router("api/RegistAPI", &account.RegistAPI{})
	beego.Router("api/CSRFToken", &account.CSRFTokenAPI{})
	beego.Router("api/UserAPI", &account.UserAPI{})
}
