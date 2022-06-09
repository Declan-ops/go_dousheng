package mapper

import (
	"go_dousheng/model"

	"sync"
)

type UserDao struct {
	// 接口类型
}

var (
	userDao     *UserDao
	userDaoOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userDaoOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryPostsByParentId(username string) *model.User {

	return userInitMap[username]
}

func (d *UserDao) UpdateMapUser(user *model.User) {
	// 更新信息
	userInitMap[user.Name] = user
}

func (d *UserDao) QueryUserByName(username string, user *model.User) {
	// 连接数据库，查询所有用户信息，放入缓存中
	db := InitDB()
	db.Where("name = ?", username).Debug().First(&user)

}

func (d *UserDao) QueryUserAttentionAndFollow(user *model.User) *model.UserVO {

	// 连接数据库，查询所有用户信息，放入缓存中
	db := InitDB()
	user_vo := model.UserVO{}
	// 粉丝数
	db.Model(&model.UserAttention{}).Where("user_id = ?", user.Id).Count(&user_vo.FollowerCount)
	// 关注数
	db.Model(&model.UserAttention{}).Where("other_id = ?", user.Id).Count(&user_vo.FollowCount)
	//fmt.Printf("%v", user_vo)
	// 组装
	user_vo.Id = user.Id
	user_vo.Name = user.Name
	return &user_vo

}
