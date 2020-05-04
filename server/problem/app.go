package problem

import (
	"server/problem/model"
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
	CreateTable(&model.Problem{}, &model.Fight{}, &model.History{})
}

func (App) name() string {
	return "ProblemApp"
}
