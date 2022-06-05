package controller

import (
	"github.com/gin-gonic/gin"
	"go_dousheng/service"
	"net/http"
)

func GetUserInfo(c *gin.Context) {

	// 获得用户信息
}

func UserLogin(c *gin.Context) {

	// 用户登录：先获取参数，然后通过用户名去数据库中查找对应用户，登录成功下发Token，登陆失败，退出
	username := c.Query("username")
	password := c.Query("password")

	user, err := service.QueryUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  err,
			"user_id":     nil,
			"token":       nil,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"user_id":     user.Id,
		"token":       user.Token,
	})

}
