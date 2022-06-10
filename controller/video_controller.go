package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_dousheng/model"
	"go_dousheng/service"
	"net/http"
	"strconv"
	"time"
)

func UpLoadFile(c *gin.Context) {

	// 认证token

	token := c.PostForm("token")
	title := c.PostForm("title")

	//f, err := c.FormFile("data") // multipart/form-data

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s"), err.Error())
	}

	videos := model.Videos{
		form,
		title,
		token,
	}
	err = service.UploadVideo(&videos)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			402,
			err.Error(),
		})
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 200,
		StatusMsg:  "上传成功！",
	})

	/*	video := VideoFile{
			user,
			title,
			f,
		}
		// 保存在指定路径
		c.SaveUploadedFile(f, "upload/"+f.Filename)
		// 将文件上传到 阿里云 oss
		UploadAliyunOss(f, &video)
	*/

}

func PublishList(c *gin.Context) {

	// 认证token
	token := c.Query("token")
	user_id := c.Query("user_id")
	userId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusConflict, model.Response{
			408,
			err.Error(),
		})
	}
	// 获取用户发布的所有视频
	user := &model.User{
		Id:    userId,
		Token: token,
	}
	video_vo_List, err := service.QueryUserVideoList(user)
	if err != nil {
		c.JSON(http.StatusConflict, model.Response{
			408,
			err.Error(),
		})
	}

	c.JSON(http.StatusOK, model.VideoVOList{
		model.Response{
			0,
			"",
		},
		video_vo_List,
	})

}
func Feed(c *gin.Context) {

	timeUnix := time.Now().Unix() //获取当前时间戳
	// int64到string
	timeunixString := strconv.FormatInt(timeUnix, 10)
	// 获取视频流
	latest_time := c.DefaultQuery("latest_time", timeunixString)
	token := c.Query("token")

	video_vo_List, err := service.VideoFeed(latest_time, token)
	if err != nil {
		c.JSON(http.StatusConflict, model.Response{
			408,
			err.Error(),
		})
	}

	c.JSON(http.StatusOK, model.VideoVOList{
		model.Response{
			0,
			"",
		},
		video_vo_List,
	})

}
