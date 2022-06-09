package mapper

import "go_dousheng/model"

func SaveVideo(video *model.Video) {

	db := InitDB()
	db.Create(video)

}
