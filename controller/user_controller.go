package controller

import (
	"github.com/gin-gonic/gin"
	"go_dousheng/model"
	"go_dousheng/service"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	// 获得用户信息

	user_id := c.Query("user_id")
	token := c.Query("token")
	userVO, err := service.QueryUserInfo(user_id, token)
	if err != nil {
		c.JSON(http.StatusOK, model.UserVO{
			Response: model.Response{
				StatusCode: 406,
				StatusMsg:  err.Error(),
			},
		})
	}

	c.JSON(http.StatusOK, model.UserResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取用户信息成功！",
		},
		User: *userVO,
	})

	/*	c.JSON(http.StatusOK, model.UserVO{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  "获取用户信息成功！",
			},
			Id:            userVO.Id,
			Name:          userVO.Name,
			FollowCount:   userVO.FollowCount,
			FollowerCount: userVO.FollowerCount,
			IsFollow:      userVO.IsFollow,
		})
	*/
}

func UserLogin(c *gin.Context) {

	// 用户登录：先获取参数，然后通过用户名去数据库中查找对应用户，登录成功下发Token，登陆失败，退出
	username := c.Query("username")
	password := c.Query("password")

	user, err := service.QueryUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "登陆成功！",
			"user_id":     nil,
			"token":       nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"user_id":     user.Id,
		"token":       user.Token,
	})

}

func Register(c *gin.Context) {
	// 用户注册
	username := c.Query("username")
	password := c.Query("password")

	user, err := service.SaveUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, model.UserPOJO{
			Response: model.Response{
				406,
				err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, model.UserPOJO{
		Response: model.Response{
			0,
			"注册成功",
		},
		UserId: user.Id,
		Token:  user.Token,
	})

}
