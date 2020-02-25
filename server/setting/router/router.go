package router

import (
	"github.com/astaxie/beego"
	account "server/account/controller"
)

func InitRouter() {
	beego.Router("api/LoginAPI", &account.LoginAPI{})
	beego.Router("api/RegistAPI", &account.RegistAPI{})
	beego.Router("api/LogoutAPI", &account.LogoutAPI{})
	beego.Router("api/CSRFTokenAPI", &account.CSRFTokenAPI{})
	beego.Router("api/UserAPI", &account.UserAPI{})
	beego.Router("api/DynamicAPI", &account.DynamicAPI{})
	beego.Router("api/CommentsAPI", &account.CommentsAPI{})
}
