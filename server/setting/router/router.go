package router

import (
	"github.com/astaxie/beego"
	account "server/account/controller"
	problem "server/problem/controller"
)

func InitRouter() {
	beego.Router("api/LoginAPI", &account.LoginAPI{})
	beego.Router("api/RegistAPI", &account.RegistAPI{})
	beego.Router("api/LogoutAPI", &account.LogoutAPI{})
	beego.Router("api/CSRFTokenAPI", &account.CSRFTokenAPI{})
	beego.Router("api/UserAPI", &account.UserAPI{})
	beego.Router("api/UserCountAPI", &account.UserCountAPI{})
	beego.Router("api/DynamicAPI", &account.DynamicAPI{})
	beego.Router("api/CommentsAPI", &account.CommentsAPI{})
	beego.Router("api/ProblemAPI", &problem.ProblemAPI{})
	beego.Router("api/ProblemCountAPI", &problem.ProblemCountAPI{})
}
