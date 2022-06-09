package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go_dousheng/mapper"
	"go_dousheng/model"
	"time"
)

var keyInfo = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"

func GetToken(user model.User) string {
	//自定义一个key

	//将部分用户信息保存到map并转换为json
	//info := map[string]model.User{}
	//info[user.Name] = user
	// 把user变成json
	dataByte, _ := json.Marshal(user)
	var dataStr = string(dataByte)

	//使用Claim保存json
	//这里是个例子，并包含了一个故意签发一个已过期的token
	data := jwt.StandardClaims{Subject: dataStr, ExpiresAt: time.Now().Unix() + 1000*360}
	tokenInfo := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	//生成token字符串
	tokenStr, _ := tokenInfo.SignedString([]byte(keyInfo))
	//fmt.Println("myToken is: ", tokenStr)
	return tokenStr

}

func GetTokenInfo(tokenStr string) (jwt.MapClaims, error) {

	//将token字符串转换为token对象（结构体更确切点吧，go很孤单，没有对象。。。）
	tokenInfo, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return keyInfo, nil
	})

	//校验错误（基本）
	err := tokenInfo.Claims.Valid()
	if err != nil {
		println(err.Error())
		return nil, err
	}

	finToken := tokenInfo.Claims.(jwt.MapClaims)

	// 检验是否过期
	succ := finToken.VerifyExpiresAt(time.Now().Unix(), true)
	if !succ {
		// 过期了提示
		return nil, errors.New("token失效!")

	}

	//获取token中保存的用户信息
	fmt.Println(finToken["sub"])
	return finToken, nil

}

func CheckToken(tokenStr string) (*model.User, error) {

	//将token字符串转换为token对象（结构体更确切点吧，go很孤单，没有对象。。。）
	tokenInfo, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return keyInfo, nil
	})

	//校验错误（基本）
	err := tokenInfo.Claims.Valid()
	if err != nil {
		println(err.Error())
		return nil, err
	}

	finToken := tokenInfo.Claims.(jwt.MapClaims)

	// 检验是否过期
	succ := finToken.VerifyExpiresAt(time.Now().Unix(), true)
	if !succ {
		// 过期了提示
		return nil, errors.New("token失效!")

	}

	var user model.User

	user_json := finToken["sub"].(string)

	error := json.Unmarshal([]byte(user_json), &user)

	if error != nil {

		return nil, nil
	}

	var user2 model.User
	// 从数据库当中根据名字取出对应的user进行对比
	mapper.NewUserDaoInstance().QueryUserByName(user.Name, &user2)
	if user2.Name != user.Name || user2.Password != user.Password {
		return nil, errors.New("Token认证失败！")
	}
	return &user2, nil

}
