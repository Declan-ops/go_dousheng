package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db = Init()

func Init() *gorm.DB {

	dsn := "root:cheng@tcp(120.25.197.142:3306)/ds_dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
