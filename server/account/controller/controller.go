package controller

import (
	"server/account/constant"
	"server/account/model"
	"server/account/serializer"
	"server/account/utils"
	"server/account/validation"
	db "server/setting/model"
	"server/setting/request"
)

type LoginAPI struct{ request.Controller }
type RegistAPI struct{ request.Controller }
type LogoutAPI struct{ request.Controller }
type UserAPI struct{ request.Controller }
type DynamicAPI struct{ request.Controller }
type CommentsAPI struct{ request.Controller }
type CSRFTokenAPI struct{ request.Controller }

func (c *LoginAPI) Post() {
	data := validation.LoginValid{}
	c.Check(&data, false)
	phoneEmail := data.PhoneEmail
	password := data.Password

	user := model.User{}
	if db.GetDB().Where("phone = ? or email = ?", phoneEmail, phoneEmail).First(&user).Error != nil {
		c.Error("账号不存在!")
		return
	}

	if user.Password != utils.Encrypt(password) {
		c.Error("密码错误!")
	}

	c.Login(user)
	c.Success(nil)
}

func (c *LogoutAPI) Post() {
	c.Check(nil, true, "all")
	c.Logout()
	c.Success(nil)
}

func (c *RegistAPI) Post() {
	data := validation.RegistValid{}
	c.Check(&data, false)
	phone := data.Phone
	password := data.Password
	username := data.Username

	if db.GetDB().Where("phone = ?", phone).First(&model.User{}).Error == nil {
		c.Error("该手机号已被注册!")
		return
	}

	user := model.User{Username: username, Password: utils.Encrypt(password), Phone: phone}
	db.GetDB().Create(&user)
	c.Success(nil)
}

func (c *UserAPI) Get() {
	c.Check(nil, true, "all")
	id, _ := c.GetInt("id")
	type User struct {
		serializer.UserSerialize
	}
	var users []User
	user := model.User{}

	if id != 0 {
		db.GetDB().Where("id = ?", id).First(&user)
		db.GetDB().Model(&user).Related(&user.UserInfo)
		c.Success(user)
		return
	}
	if c.RquestUser().UserType != constant.Admin {
		c.Error("权限不足")
		return
	}
	db.GetDB().Preload("UserInfo").Find(&users)
	c.Success(users)
}

func (c *UserAPI) Put() {
	data := validation.UpdateUserValid{}
	c.Check(&data, true, "all")
	id := data.Id
	phone := data.Phone
	username := data.Username
	email := data.Email
	password := data.Password
	declaration := data.Declaration
	avatar := data.Avatar

	user := model.User{}
	userInfo := model.UserInfo{}

	if db.GetDB().Where("id = ?", id).First(&model.User{}).Error == nil {

		db.GetDB().Where("id = ?", id).Model(&user).Updates(model.User{Phone: phone, Username: username, Email: email, Password: password})
		db.GetDB().Where("user_id = ?", id).Model(&userInfo).Updates(model.UserInfo{Declaration: declaration, Avatar: avatar})

	} else {
		c.Error("系统没有该人员")
		return
	}

	c.Success(nil)
}

func (c *UserAPI) Delete() {
	data := validation.DeleteByIdValid{}
	c.Check(&data, true, "admin")
	id := data.Id
	if db.GetDB().Where("id = ?", id).First(&model.User{}).Error == nil {
		db.GetDB().Delete(model.User{}, "id = ?", id)
	} else {
		c.Error("系统没有该人员")
		return
	}
	c.Success(nil)
}

func (c *DynamicAPI) Get() {

	id, _ := c.GetInt("id")
	type Dynamic struct {
		serializer.DynamicSerialize
	}
	var dynamics []Dynamic
	if id != 0 {
		db.GetDB().Where("user_id = ?", id).Find(&dynamics)

		c.Success(&dynamics)
		return
	}
	db.GetDB().Find(&dynamics)
	c.Success(&dynamics)
}

func (c *DynamicAPI) Post() {

	data := validation.DynamicValid{}
	c.Check(&data, true)

	userId := c.RquestUser().ID
	content := data.Content

	imgPath := data.ImgPath
	imgPathUUID := utils.ImgToUUID(imgPath)

	if db.GetDB().Where("id = ?", userId).First(&model.User{}).Error != nil {

		c.Error("系统没有该人员")
		return
	}

	dynamic := model.Dynamic{UserID: userId, Content: content, ImgPath: imgPathUUID}
	db.GetDB().Create(&dynamic)
	c.Success(nil)
}

func (c *DynamicAPI) Delete() {

	data := validation.DeleteByIdValid{}
	c.Check(&data, true, "all")

	id := data.Id

	if db.GetDB().Where("id = ?", id).First(&model.Dynamic{}).Error == nil {
		db.GetDB().Delete(model.Dynamic{}, "id = ?", id)
	} else {
		c.Error("系统没有该条动态")
		return
	}
	c.Success(nil)

}

func (c *CommentsAPI) GET() {
	id, _ := c.GetInt("id")
	dynamicId, _ := c.GetInt("dynamic_id")
	userId, _ := c.GetInt("comments_id")

	var users []serializer.CommentsSerialize

	if id != 0 {
		user := serializer.CommentsSerialize{}
		db.GetDB().Where("id = ?", id).First(&user, "User")
		c.Success(user)
		return
	} else if dynamicId != 0 {
		db.GetDB().Where("dynamic_id = ?", dynamicId).First(&users, "User")
	} else if userId != 0 {
		db.GetDB().Where("user_id = ?", userId).First(&users, "User")
	} else {
		db.GetDB().Find(&users, "User")

	}
	c.Success(users)
}

func (c *CommentsAPI) POST() {
	data := validation.CommentsValid{}
	c.Check(&data, false)

	userId := data.UserId
	dynamicId := data.DynamicId
	commentsId := data.CommentsId
	content := data.Content
	imgPath := data.ImgPath

	if db.GetDB().Where("id = ?", userId).First(&model.User{}).Error != nil {
		c.Error("系统没有该人员")
		return
	}

	comment := model.Comment{UserId: userId, DynamicId: dynamicId, CommentsId: commentsId, Content: content, ImgPath: imgPath}
	db.GetDB().Create(&comment)
	c.Success(nil)

}

func (c *CSRFTokenAPI) Get() {
	c.Success(map[string]string{"Csrftoken": c.XSRFToken()})
}
