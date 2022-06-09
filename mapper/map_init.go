package mapper

import (
	"fmt"
	"go_dousheng/model"
)

var (
	// 私有
	userInitMap map[string]*model.User
	//postInitMap  map[int64][]*Post
)

func IntitUserMap() error {

	// 连接数据库，查询所有用户信息，放入缓存中
	db := InitDB()
	data := make([]*model.User, 0)
	err := db.Find(&data).Error
	if err != nil {
		fmt.Println("查询失败，错误：", err)
		return err
	}

	userTempMap := make(map[string]*model.User)
	// 遍历数据，放入map中
	for _, v := range data {

		userTempMap[v.Name] = v
		fmt.Println("用户：", userInitMap[v.Token])
	}
	userInitMap = userTempMap
	return nil

}
