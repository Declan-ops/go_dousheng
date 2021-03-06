package router

import (
	"github.com/gin-gonic/gin"
	"go_dousheng/controller"
)

func InitRouter(r *gin.Engine) {

	// public directory is used to serve static resources
	r.Static("/static", "./public")
	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", controller.GetUserInfo)
	apiRouter.POST("/user/login/", controller.UserLogin)
	apiRouter.POST("/publish/action/", controller.UpLoadFile)
	apiRouter.GET("/publish/list/", controller.PublishList)
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)

	/*
		// extra apis - I
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		apiRouter.GET("/favorite/list/", controller.FavoriteList)
		apiRouter.POST("/comment/action/", controller.CommentAction)
		apiRouter.GET("/comment/list/", controller.CommentList)

		// extra apis - II
		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.FollowList)
		apiRouter.GET("/relation/follower/list/", controller.FollowerList)*/
}
