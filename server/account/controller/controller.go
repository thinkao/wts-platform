package controller

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"server/account/constant"
	"server/account/model"
	"server/account/serializer"
	"server/account/utils"
	"server/account/validation"
	"server/setting/config"
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
	if c.RequestUser().UserType != constant.Admin {
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
	fmt.Println("----",data.Id)
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

	c.Check(nil, true, "all")
	id, _ := c.GetInt("id")
	type Dynamic struct {
		serializer.DynamicSerialize
	}

	/*select avatar,username,content,img_path,d.create_time from user as u,user_info as i,dynamic as d where u.id = i.user_id and u.id = d.user_id order by d.create_time desc;*/

	var results []model.Result
	if id != 0 {
		db.GetDB().Table("user").Select("user.username,user_info.avatar,dynamic.content,dynamic.img_path,dynamic.create_time").Joins("left join user_info on user.id = user_info.user_id").Joins("right join dynamic on dynamic.user_id = user.id").Where("user.id = ?",id).Order("dynamic.create_time desc").Scan(&results)

	}else {
		db.GetDB().Table("user").Select("user.username,user_info.avatar,dynamic.content,dynamic.img_path,dynamic.create_time").Joins("left join user_info on user.id = user_info.user_id").Joins("right join dynamic on dynamic.user_id = user.id").Order("dynamic.create_time desc").Scan(&results)
	}

	for i:=0;i< len(results);i++ {
		results[i].ImgPath = config.IMAGE_FILE_PATH + results[i].ImgPath + ".png"
	}

	c.Success(&results)
}


func (c *DynamicAPI) Post() {

	data := validation.DynamicValid{}
	c.Check(&data, true,"all")

	userId := c.RequestUser().ID

	content := data.Content
	imgPath := data.ImgPath

	var img_file_path string

	f, _, _ := c.GetFile("file")
	if f != nil {
		u2 := uuid.NewV4()
		img_file_path = config.IMAGE_FILE_PATH + u2.String() + ".png"
		c.SaveToFile("file", img_file_path)
		img := model.Dynamic{
			ImgPath:   u2.String(),
		}
		c.Success(img)
		return
	}


	if db.GetDB().Where("id = ?", userId).First(&model.User{}).Error != nil {
		c.Error("系统没有该人员")
		return
	}

	if content == "" && imgPath == ""{
		c.Success(nil)
		return
	}

	dynamic := model.Dynamic{UserID: userId,Content:content,ImgPath:imgPath}
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
	c.Success(&users)
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
