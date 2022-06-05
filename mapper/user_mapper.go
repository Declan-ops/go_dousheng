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
