package request

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	validator "gopkg.in/go-playground/validator.v9"
	"reflect"
	"server/account/model"
	"strconv"
)

type Controller struct {
	beego.Controller
}

type ResponseData struct {
	Err  interface{} `json:"err"`
	Data interface{} `json:"data"`
}

func (c *Controller) Response(err interface{}, data interface{}, code int) {
	/*u.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", u.Ctx.Request.Header.Get("Origin"))*/
	c.Ctx.ResponseWriter.WriteHeader(code)
	resp := ResponseData{Err: err, Data: data}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *Controller) Success(data interface{}) {
	c.Response(nil, data, 200)
}

func (c *Controller) Error(msg interface{}) {
	data := map[string]interface{}{"message": msg}
	c.Response("err", data, 200)
}

func (c *Controller) SpecificError(errType string, message interface{}) {
	data := map[string]interface{}{"message": message}
	c.Response(errType, data, 200)
	c.StopRun()
}

func (c *Controller) ServerError() {
	c.SpecificError("server-error", "服务器错误！")
}

func (c *Controller) XSRFError(msg string) {
	c.SpecificError("xsrf-error", msg)
}

func (c *Controller) PermissionError() {
	c.SpecificError("permission-denied", "访问受限！")
}

func (c *Controller) LoginRequiredError() {
	c.SpecificError("login-required", "需要登录！")
}

func (c *Controller) InvalidDataError(msg interface{}) {
	c.SpecificError("invalid-data", msg)
}

func (c *Controller) Check(validRequest interface{}, loginRequired bool, permissions ...string) {
	if loginRequired == true {
		if permissions == nil {
			panic("Invalid check condition")
		}
		session := c.GetSession("user")
		if session == nil {
			c.LoginRequiredError()
		}
		hasPermission := false
		for _, permission := range permissions {
			if permission == "all" {
				hasPermission = true
				break
			}

			user := reflect.ValueOf(session)
			fmt.Printf("**************",user)
			if user.FieldByName("UserType").String() == permission {
				hasPermission = true
				break
			}
		}

		if hasPermission == false {
			c.PermissionError()
		}
	}

	if validRequest == nil {
		return
	}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, validRequest)

	validate := validator.New()
	errs := validate.Struct(validRequest)

	if errs != nil {

		msg := make(map[string]string)
		for _, err := range errs.(validator.ValidationErrors) {
			msg[err.Field()] = err.Tag()
		}
		c.InvalidDataError(msg)
	}

}

func (c *Controller) Login(user model.User) {
	c.SetSession("user", user)
}

func (c *Controller) Logout() {
	c.DelSession("user")
}

func (c *Controller) RquestUser() model.User{
	session := c.GetSession("user")
	reflectUser := reflect.ValueOf(session)
	id, _ := strconv.Atoi(reflectUser.FieldByName("Id").String())
	user := model.User{
		ID: id,
		Username: reflectUser.FieldByName("Username").String(),
		Password: reflectUser.FieldByName("Password").String(),
		UserType: reflectUser.FieldByName("UserType").String(),
		Phone: reflectUser.FieldByName("Phone").String(),
		Email: reflectUser.FieldByName("Email").String(),
	}
	return user
}

func (c *Controller) CheckXSRFCookie() bool {
	if !c.EnableXSRF {
		return true
	}
	token := c.Ctx.Input.Query("_xsrf")
	if token == "" {
		token = c.Ctx.Request.Header.Get("X-Xsrftoken")
	}
	if token == "" {
		token = c.Ctx.Request.Header.Get("X-Csrftoken")
	}
	if token == "" {
		c.XSRFError("'_xsrf' argument missing from cookie")
	}
	if c.XSRFToken() != token {
		c.XSRFError("XSRF cookie does not match")
	}
	return true
}
