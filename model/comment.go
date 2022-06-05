package model

import "time"

type Comment struct {
	id          int64     `json:"id,omitempty"`
	videoId     int64     `json:"videoId,omitempty"`
	userId      int64     `json:"userId,omitempty"`
	content     string    `json:"id,omitempty"`
	create_time time.Time `json:"create_Time,omitempty"`
}

func (Comment) TableName() string {
	return "comment_table"
}
