package util

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go_dousheng/model"
	"time"
)

var keyInfo = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"

func GetToken(user model.User) string {
	//自定义一个key

	//将部分用户信息保存到map并转换为json
	info := map[string]model.User{}
	info[user.Name] = user
	dataByte, _ := json.Marshal(info)
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
	//获取token中保存的用户信息
	fmt.Println(finToken["sub"])
	return finToken, nil

}
