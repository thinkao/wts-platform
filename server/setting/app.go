package setting

import (
	"server/setting/config"
	_ "server/setting/model"
	"server/setting/router"
)

type App struct{}

func (App) SetUp() {
	config.InitConfig()
	router.InitRouter()
}

func (App) Name() string {
	return "AccountApp"
}
