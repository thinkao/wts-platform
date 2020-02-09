package main

import (
	"github.com/astaxie/beego"
	"server/account"
	"server/problem"
	"server/setting"
)

func main() {
	setting.App{}.SetUp()
	account.App{}.SetUp()
	problem.App{}.SetUp()
	beego.Run()
}
