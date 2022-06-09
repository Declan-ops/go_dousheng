package model

import "mime/multipart"

type Video struct {
	Id         int64
	AuthorId   int64
	PlayUrl    string
	CoverUrl   string
	Title      string
	CreateTime int64
}

func (Video) TableName() string {
	return "video_table"
}

type VideoUserFavorite struct {
	Id      int64 `gorm:"id"`
	VideoId int64 `gorm:"video_id"`
	UserId  int64 `gorm:"user_id"`
}

func (VideoUserFavorite) TableName() string {
	return "video_user_favorite_table"
}

type Videos struct {
	Data  *multipart.Form
	Title string
	Token string
}
