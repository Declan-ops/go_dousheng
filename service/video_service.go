package service

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go_dousheng/mapper"
	"go_dousheng/model"
	"go_dousheng/util"
	"mime/multipart"
	"os"
	"time"
)

type VideoFile struct {
	user  *model.User
	title string
	file  *multipart.FileHeader
}

func UploadVideo(video *model.Videos) error {

	user, err2 := util.CheckToken(video.Token)
	if err2 != nil && user != nil {

		if user == nil {
			return errors.New("token认证失败！未查询到该用户")
		} else {
			return errors.New(err2.Error())
		}

	}
	// 支持多文件
	f := video.Data.File["data"]
	for _, file := range f {
		video := VideoFile{
			user,
			video.Title,
			file,
		}
		// 保存在指定路径
		// c.SaveUploadedFile(file, "upload/"+file.Filename)
		// 将文件上传到 阿里云 oss
		real_path := UploadAliyunOss(file, &video)
		// 保存到数据库中
		cover_url := real_path + "?x-oss-process=video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast"
		timeUnix := time.Now().Unix() //获取当前时间戳

		var video_save = &model.Video{
			AuthorId:   user.Id,
			PlayUrl:    real_path,
			CoverUrl:   cover_url,
			Title:      video.title,
			CreateTime: timeUnix,
		}
		mapper.NewVideoDaoInstance().SaveVideo(video_save)

	}
	return nil
}

// 将文件上传到 阿里云 oss
func UploadAliyunOss(file *multipart.FileHeader, video *VideoFile) string {

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
	real_path := "https://jxau7124.oss-cn-shenzhen.aliyuncs.com/" + path
	return real_path
}

func QueryUserVideoList(user *model.User) ([]*model.VideoVO, error) {

	// 解析token，获取当前用户
	db_user, err2 := util.CheckToken(user.Token)
	if err2 != nil && db_user != nil && db_user.Id == user.Id {

		if user == nil {
			return nil, errors.New("token认证失败！未查询到该用户")
		} else {
			return nil, errors.New(err2.Error())
		}

	}

	// 查询数据库，获取发布视频列表
	video_list, err := mapper.NewVideoDaoInstance().QueryUserVideo(db_user)
	if err != nil {
		return nil, err
	}
	return video_list, nil
}

func VideoFeed(latest_time, token string) ([]*model.VideoVO, error) {

	// 查询数据库，获取所有发布视频列表
	video_list, err := mapper.NewVideoDaoInstance().QueryAllVideos(latest_time)
	if err != nil {
		return nil, err
	}
	return video_list, nil
}
