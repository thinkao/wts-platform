package controller

import (
	"server/account/model"
	"server/account/serializer"
	"server/account/utils"
	"server/account/validation"
	db "server/setting/model"
	"server/setting/request"
)

type LoginAPI struct{ request.Controller }
type RegistAPI struct{ request.Controller }
type UserAPI struct{ request.Controller }
type CSRFTokenAPI struct{ request.Controller }

func (c *LoginAPI) Post() {
	data := validation.LoginValid{}
	c.Check(&data, false)
	usernamePhone := data.UsernamePhone
	password := data.Password

	user := model.User{}
	err := db.GetDB().Where("username = ?", usernamePhone).First(&user).Error
	if err != nil {
		err := db.GetDB().Where("phone = ?", usernamePhone).First(&user).Error
		if err != nil {
			c.Error("账号不存在!")
			return
		}
	}

	if user.Password != utils.Encrypt(password) {
		c.Error("密码错误!")
	}
	c.SetSession("user", user)
	c.Success(nil)
}

func (c *RegistAPI) Post() {
	data := validation.RegistValid{}
	c.Check(&data, false)
	username := data.Username
	password := data.Password
	phone := data.Phone

	if db.GetDB().Where("username = ?", username).First(&model.User{}).Error == nil {
		c.Error("用户名已存在!")
		return
	}

	if db.GetDB().Where("phone = ?", phone).First(&model.User{}).Error == nil {
		c.Error("该手机号已被注册!")
		return
	}

	user := model.User{Username: username, Password: utils.Encrypt(password), Phone: phone}
	db.GetDB().Create(&user)
	c.Success(nil)
}

func (c *UserAPI) Get() {
	//c.Check(nil, true, "all")
	type User struct {
		serializer.UserSerialize
	}
	var users []User
	db.GetDB().Find(&users)
	c.Success(users)
}

func (c *UserAPI) Post() {
	c.Error(nil)
}

func (c *UserAPI) Put() {
	c.Success(nil)
}

func (c *UserAPI) Delete() {
	c.Success(nil)
}

func (c *CSRFTokenAPI) Get() {
	c.Success(map[string]string{"X-Csrftoken": c.XSRFToken()})
}
