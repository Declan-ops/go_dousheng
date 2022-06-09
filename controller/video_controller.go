package controller

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go_dousheng/model"
	"go_dousheng/util"
	"mime/multipart"
	"net/http"
	"os"
)

type VideoFile struct {
	user  *model.User
	title string
	file  *multipart.FileHeader
}

func UpLoadFile(c *gin.Context) {

	// 认证token

	token := c.PostForm("token")
	title := c.PostForm("title")
	f, err := c.FormFile("data") // multipart/form-data

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user, err2 := util.CheckToken(token)
	if err2 != nil && user != nil {

		if user == nil {
			c.JSON(http.StatusBadRequest, model.Response{
				402,
				"token认证失败！未查询到该用户",
			})
		} else {
			c.JSON(http.StatusBadRequest, model.Response{
				402,
				err2.Error(),
			})
		}

		return

	}
	video := VideoFile{
		user,
		title,
		f,
	}
	// 保存在指定路径
	c.SaveUploadedFile(f, "upload/"+f.Filename)

	// 将文件上传到 阿里云 oss
	UploadAliyunOss(f, &video)
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 200,
		StatusMsg:  "上传成功！",
	})
}

// 将文件上传到 阿里云 oss
func UploadAliyunOss(file *multipart.FileHeader, video *VideoFile) {

	Endpoint := "https://cn-shenzhen.oss.aliyuncs.com"  // 这里的是广州区，
	AccessKeyID := "LTAI5tMBdX1b5HckRZZFxrKt"           //
	AccessKeySecret := "K3BvMYNTnywgdf2kKoQmtniZflXUNl" //

	client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 指定bucket
	bucket, err := client.Bucket("jxau7124") // 根据自己的填写
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	defer src.Close()

	// 将文件流上传至test目录下
	path := video.user.Name + "/" + video.title + file.Filename
	err = bucket.PutObject(path, src)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fmt.Println("file upload success")
}
