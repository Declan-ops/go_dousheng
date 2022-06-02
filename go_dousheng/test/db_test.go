package test

import (
	"fmt"
	"go_dousheng/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormTest(t *testing.T) {
	dsn := "root:cheng@tcp(120.25.197.142:3306)/ds_dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("连接失败", err)
	}
	// 查询数据
	// 插入数据
	u1 := model.Video{
		1,
		1,
		"https://www.w3schools.com/html/movie.mp4",
		"https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		"小熊洗澡",
		1654179366,
	}
	if err := db.Create(&u1).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
	data := make([]*model.Video, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", data)
	fmt.Println("1111")
	for _, v := range data {
		fmt.Println("1111")
		fmt.Printf("%v\n", v)
	}
}
