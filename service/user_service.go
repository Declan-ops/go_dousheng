package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"go_dousheng/mapper"
	"go_dousheng/model"
	"go_dousheng/util"
	"strconv"
	"strings"
)

func QueryUser(username, password string) (*model.User, error) {

	user := mapper.NewUserDaoInstance().QueryPostsByParentId(username)

	if user == nil {
		return nil, errors.New("用户不存在！请检查用户名")
	}

	// 进行md5加密
	newSig := md5.Sum([]byte(password)) //转成加密编码
	// 将编码转换为字符串
	newArr := fmt.Sprintf("%x", newSig)
	//输出字符串字母都是小写，转换为大写
	password = strings.ToTitle(newArr)

	if password == user.Password {

		a := false
		// 如果没有token就下发token
		token := user.Token
		if len(token) == 0 {
			// 下发token
			a = true
			user.Token = util.GetToken(*user)
		}

		// 检查token是否过期，过期就更新
		_, err := util.GetTokenInfo(user.Token)
		if err != nil {

			if err.Error() == "token失效!" {
				// token过期，更新token
				a = true
				user.Token = util.GetToken(*user)
			} else {
				return nil, err
			}
		}

		if a {
			// token更新，保存进数据库
			go func() {
				db := mapper.InitDB()
				db.Model(&user).Update("token", user.Token)
			}()

			// 缓存也得更新
			mapper.NewUserDaoInstance().UpdateMapUser(user)

		}

		return user, nil
	} else {
		return nil, errors.New("用户密码错误！")
	}

	return nil, nil
}

func QueryUserInfo(user_id, token string) (*model.UserVO, error) {

	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		return nil, errors.New("程序异常")
	}

	// 进行token认证
	user, err := util.CheckToken(token)
	if err != nil || user.Id != userId {
		return nil, errors.New("token认证失败！")
	}

	// 根据条件查找用户的粉丝数以及关注数
	userVO := mapper.NewUserDaoInstance().QueryUserAttentionAndFollow(user)
	return userVO, nil

}

func SaveUser(username, password string) (*model.User, error) {

	password = util.MD5(password)
	user := model.User{
		Name:     username,
		Password: password,
	}
	user.Token = util.GetToken(user)
	err := mapper.NewUserDaoInstance().SaveUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
