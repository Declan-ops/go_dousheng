package model

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// 直接返回给浏览器的实体模型
type VideoVO struct {
	Response      Response
	Author        UserVO
	Id            int64  `json:"id,omitempty"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommenetCount int64  `json:"commenet_count,omitempty"`
	Title         string `json:"title,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	NextTime      int64  `json:"next_time,omitempty"` // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type UserVO struct {
	Response      Response
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`   // 关注总数
	FollowerCount int64  `json:"follower_count,omitempty"` // 粉丝总数
	IsFollow      bool   `json:"is_follow,omitempty"`
}
