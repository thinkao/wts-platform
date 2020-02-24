package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"server/setting/config"
	"time"
)

var serverDB *gorm.DB

func init() {
	var (
		db  *gorm.DB
		err error
	)
	for i := 0; i < 5; i++ {
		db, err = gorm.Open("mysql", config.DbAddr)
		if err != nil {
			//log.Errorf("db connect failed: %s", err)
			time.Sleep(3 * time.Second)
			//log.Infof("retry to connect db")
		} else {
			err = nil
			break
		}
	}
	if err != nil {
		panic(err)
	}

	serverDB = db
	db.SingularTable(true)
	/*打印原生SQL*/
	db.LogMode(true)
	serverDB.DB().SetConnMaxLifetime(100)
	serverDB.DB().SetMaxIdleConns(20)
}
func GetDB() *gorm.DB {
	return serverDB
}
