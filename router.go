package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/comment"
	"github.com/zheng-yi-yi/simpledouyin/controllers/favorite"
	"github.com/zheng-yi-yi/simpledouyin/controllers/message"
	"github.com/zheng-yi-yi/simpledouyin/controllers/relation"
	"github.com/zheng-yi-yi/simpledouyin/controllers/user"
	"github.com/zheng-yi-yi/simpledouyin/controllers/video"
	"github.com/zheng-yi-yi/simpledouyin/middlewares"
)

// route initialization function
func initRouter(r *gin.Engine) {
	// basicSetup
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	apiRouter := r.Group("/douyin")
	// fundamental features：
	apiRouter.GET("/feed/", video.Feed)                                    //  视频流
	apiRouter.POST("/user/register/", user.Register)                       // 用户注册
	apiRouter.POST("/user/login/", user.Login)                             // 用户登录
	apiRouter.GET("/user/", middlewares.Auth(), user.UserInfo)             // 用户信息
	apiRouter.POST("/publish/action/", middlewares.Auth(), video.Publish)  // 视频投稿
	apiRouter.GET("/publish/list/", middlewares.Auth(), video.PublishList) // 发布列表
	// Extended Feature 1: Interactivity
	apiRouter.POST("/favorite/action/", middlewares.Auth(), favorite.FavoriteAction) // 点赞操作
	apiRouter.GET("/favorite/list/", middlewares.Auth(), favorite.FavoriteList)      // 喜欢列表
	apiRouter.POST("/comment/action/", middlewares.Auth(), comment.CommentAction)    // 评论操作
	apiRouter.GET("/comment/list/", middlewares.Auth(), comment.CommentList)         // 评论列表
	// Extended Feature 2: Social
	apiRouter.POST("/relation/action/", middlewares.Auth(), relation.RelationAction)     // 关注操作
	apiRouter.GET("/relation/follow/list/", middlewares.Auth(), relation.FollowList)     // 关注列表
	apiRouter.GET("/relation/follower/list/", middlewares.Auth(), relation.FollowerList) // 粉丝列表
	apiRouter.GET("/relation/friend/list/", middlewares.Auth(), relation.FriendList)     // 好友列表
	apiRouter.POST("/message/action/", middlewares.Auth(), message.MessageAction)        // 发送消息
	apiRouter.GET("/message/chat/", middlewares.Auth(), message.MessageChat)             // 聊天记录
}
