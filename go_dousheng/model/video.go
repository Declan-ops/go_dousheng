package model

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
