package mapper

import (
	"go_dousheng/model"
	"sync"
)

type VideoDao struct {
	// 接口类型
}

var (
	videodao     *VideoDao
	videoDaoOnce sync.Once
)

func NewVideoDaoInstance() *VideoDao {
	videoDaoOnce.Do(
		func() {
			videodao = &VideoDao{}
		})
	return videodao
}

func (*VideoDao) SaveVideo(video *model.Video) {

	db := InitDB()
	db.Create(video)

}

func (*VideoDao) QueryUserVideo(user *model.User) ([]*model.VideoVO, error) {

	db := InitDB()
	user_vo := NewUserDaoInstance().QueryUserAttentionAndFollow(user)
	videos := make([]*model.Video, 0)

	video_list := make([]*model.VideoVO, 0)
	err := db.Where("author_id = ?", user.Id).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	for _, video := range videos {

		// 组装Video_list
		video_vo := model.VideoVO{
			Author:   *user_vo,
			Id:       video.Id,
			PlayUrl:  video.PlayUrl,
			CoverUrl: video.CoverUrl,
			Title:    video.Title,
		}
		var flag int64
		// 三次查询，根据视频id寻找点赞数和评论数
		// 点赞数
		db.Model(&model.VideoUserFavorite{}).Where("video_id = ?", video.Id).Count(&video_vo.FavoriteCount)
		// 评论数
		db.Model(&model.Comment{}).Where("video_id = ?", video.Id).Count(&video_vo.CommenetCount)
		// 用户是否对改视频点过赞
		// 是否点赞
		db.Model(&model.VideoUserFavorite{}).Where("video_id = ? and user_id = ?", video.Id, user.Id).Count(&flag)
		if flag != 0 {
			video_vo.IsFavorite = true
		}
		video_list = append(video_list, &video_vo)
	}

	return video_list, nil

}

func (*VideoDao) QueryAllVideos(latest_time string) ([]*model.VideoVO, error) {

	db := InitDB()
	videos := make([]*model.Video, 0)
	video_list := make([]*model.VideoVO, 0)
	err := db.Where("create_time < ?", latest_time).Find(&videos).Limit(30).Error
	if err != nil {
		return nil, err
	}

	for _, video := range videos {

		// 组装Video_list
		video_vo := model.VideoVO{
			Id:       video.Id,
			PlayUrl:  video.PlayUrl,
			CoverUrl: video.CoverUrl,
			Title:    video.Title,
		}
		video_vo.Author.Id = video.AuthorId
		// 四次查询，根据视频id寻找点赞数和评论数，根据用户id查用户
		NewUserDaoInstance().QueryUserVOAttentionAndFollow(&video_vo.Author)
		var flag int64
		// 点赞数
		db.Model(&model.VideoUserFavorite{}).Where("video_id = ?", video.Id).Count(&video_vo.FavoriteCount)
		// 评论数
		db.Model(&model.Comment{}).Where("video_id = ?", video.Id).Count(&video_vo.CommenetCount)
		// 用户是否对改视频点过赞
		// 是否点赞
		db.Model(&model.VideoUserFavorite{}).Where("video_id = ? and user_id = ?", video.Id, video_vo.Author.Id).Count(&flag)
		if flag != 0 {
			video_vo.IsFavorite = true
		}
		video_list = append(video_list, &video_vo)
	}

	return video_list, nil
}
