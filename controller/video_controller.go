package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_dousheng/model"
	"go_dousheng/service"
	"net/http"
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
