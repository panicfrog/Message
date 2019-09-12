package dbOps;

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"message/data"
	"message/internel"
)

var DB *gorm.DB

func ConfigDB() {
	var err error
	diverName := "mysql"
	dataSource := fmt.Sprintf("%s:%s@(%s:%d)/message?charset=utf8&parseTime=True&loc=Local",
		internel.Configuration.DBUserName,
		internel.Configuration.DBPasswd,
		internel.Configuration.DBHost,
		internel.Configuration.DBPort)
	DB, err = gorm.Open(diverName, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	DB.LogMode(true)
	DB.AutoMigrate(&data.User{}, &data.Message{}, &data.Room{}, &data.Friend{}, &data.Invitation{},)
}
