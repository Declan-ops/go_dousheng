package test

import (
	"crypto/md5"
	"fmt"
	"go_dousheng/mapper"
	"go_dousheng/model"
	"go_dousheng/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"testing"
	"time"
)

func TestGormTest(t *testing.T) {
	dsn := "root:cheng@tcp(120.25.197.142:3306)/ds_dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("连接失败", err)
	}
	// 查询数据

	// 插入数据
	u1 := model.Video{
		1,
		1,
		"https://www.w3schools.com/html/movie.mp4",
		"https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		"小熊洗澡",
		1654179366,
	}
	if err := db.Create(&u1).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
	data := make([]*model.Video, 0)

	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", data)
	fmt.Println("1111")
	for _, v := range data {
		fmt.Println("1111")
		fmt.Printf("%v\n", v)
	}
}

func TestInitMap(t *testing.T) {
	// 连接数据库，查询所有用户信息，放入缓存中
	db := mapper.InitDB()

	user1 := model.User{
		1,
		"张三",
		"123",
		"123",
	}
	var user model.User
	db.Where("id = ?", user1.Id).Debug().First(&user)

	fmt.Println(user)

	return
}

func TestInitMap1(t *testing.T) {
	// 连接数据库，查询所有用户信息，放入缓存中
	db := mapper.InitDB()

	var user model.User
	db.Where("id = ?", 2).Debug().First(&user)

	fmt.Println(user.Token, "111")

	return
}

func Test2(t *testing.T) {
	// 进行md5加密
	newSig := md5.Sum([]byte("123456")) //转成加密编码
	// 将编码转换为字符串
	newArr := fmt.Sprintf("%x", newSig)
	//输出字符串字母都是小写，转换为大写
	password := strings.ToTitle(newArr)
	fmt.Println(password + "111")
}
func Test3(t *testing.T) {
	// 连接数据库，查询所有用户信息，放入缓存中

	db := mapper.InitDB()

	user_vo := model.UserVO{}
	// 粉丝数
	db.Model(&model.UserAttention{}).Where("user_id = ?", 1).Count(&user_vo.FollowerCount)
	// 关注数
	db.Model(&model.UserAttention{}).Where("other_id = ?", 1).Count(&user_vo.FollowCount)
	fmt.Printf("%v", user_vo)

}
func Test4(t *testing.T) {
	// 连接数据库，创建表

	db := mapper.InitDB()
	db.AutoMigrate(&model.VideoUserFavorite{})
}

func Test5(t *testing.T) {
	// 增加数据
	timeUnix := time.Now().Unix() //获取当前时间戳
	fmt.Println(timeUnix)
	video_save := &model.Video{

		AuthorId:   1,
		PlayUrl:    "https://jxau7124.oss-cn-shenzhen.aliyuncs.com/%E5%BC%A0%E4%B8%89/%E5%96%9C%E6%AC%A2%E4%BD%A0%E5%91%80%EF%BC%81%EF%BC%81%EF%BC%81VIDEO_20220609_150310888.mp4",
		CoverUrl:   "https://jxau7124.oss-cn-shenzhen.aliyuncs.com/%E5%BC%A0%E4%B8%89/%E5%96%9C%E6%AC%A2%E4%BD%A0%E5%91%80%EF%BC%81%EF%BC%81%EF%BC%81VIDEO_20220609_150310888.mp4?x-oss-process=video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast",
		Title:      "桌面分享",
		CreateTime: timeUnix,
	}
	db := mapper.InitDB()
	db.Debug().Create(video_save)

}
func Test6(t *testing.T) {

	password := util.MD5("123456")
	user := &model.User{
		Name:     "beixiaofeng",
		Password: password,
		Token:    "",
	}
	// 连接数据库
	db := mapper.InitDB()
	db.Create(user).Debug()
	fmt.Println(user)
}
