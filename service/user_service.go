package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"go_dousheng/mapper"
	"go_dousheng/model"
	"go_dousheng/util"
	"strings"
	"time"
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

		// 检查token是否过期，过期就更新
		finToken, erl := util.GetTokenInfo(user.Token)
		if erl != nil {
			return nil, erl
		}
		//校验下token是否过期
		succ := finToken.VerifyExpiresAt(time.Now().Unix(), true)
		if !succ {

			// 更新token
			user.Token = util.GetToken(*user)

			// 保存进数据库
			go func() {
				db := mapper.InitDB()
				db.Model(&user).Update("token", user.Token)
			}()
		}

		return user, nil
	} else {
		return nil, errors.New("用户密码错误！")
	}

	return nil, nil
}
