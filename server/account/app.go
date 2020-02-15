package account

import (
	"server/account/model"
	setting "server/setting/model"
)

type App struct{}

func CreateTable(tables ...interface{}) {
	for _, v := range tables {
		if setting.GetDB().HasTable(v) == false {
			setting.GetDB().CreateTable(v)
		}
	}
}
func (App) SetUp() {
	CreateTable(&model.User{}, &model.UserInfo{}, &model.Dynamic{}, &model.Comment{})
}

func (App) Name() string {
	return "AccountApp"
}
