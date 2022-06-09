package model

type User struct {
	/*Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`*/

	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string
}

func (User) TableName() string {
	return "user_table"
}

type UserAttention struct {
	Id      int64 `gorm:"id"`
	UserId  int64 `gorm:"userId"`
	OtherId int64 `gorm:"otherId"`
}

func (UserAttention) TableName() string {
	return "user_attention_table"
}
