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

// 路由初始化函数
func initRouter(r *gin.Engine) {
	// 基本设置
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	apiRouter := r.Group("/douyin")
	// （1）基础功能：
	apiRouter.GET("/feed/", video.Feed)                                              //  视频流
	apiRouter.POST("/user/register/", user.Register)                                 // 用户注册
	apiRouter.POST("/user/login/", user.Login)                                       // 用户登录
	apiRouter.GET("/user/", middlewares.JWTAuth("query"), user.UserInfo)             // 用户信息
	apiRouter.POST("/publish/action/", middlewares.JWTAuth("form"), video.Publish)   // 视频投稿
	apiRouter.GET("/publish/list/", middlewares.JWTAuth("query"), video.PublishList) // 发布列表
	// （2）互动功能：
	apiRouter.POST("/favorite/action/", middlewares.JWTAuth("query"), favorite.FavoriteAction) // 点赞操作
	apiRouter.GET("/favorite/list/", middlewares.JWTAuth("query"), favorite.FavoriteList)      // 喜欢列表
	apiRouter.POST("/comment/action/", middlewares.JWTAuth("query"), comment.CommentAction)    // 评论操作
	apiRouter.GET("/comment/list/", comment.CommentList)                                       // 评论列表
	// （3）社交功能：
	apiRouter.POST("/relation/action/", middlewares.JWTAuth("query"), relation.RelationAction)     // 关注操作
	apiRouter.GET("/relation/follow/list/", middlewares.JWTAuth("query"), relation.FollowList)     // 关注列表
	apiRouter.GET("/relation/follower/list/", middlewares.JWTAuth("query"), relation.FollowerList) // 粉丝列表
	apiRouter.GET("/relation/friend/list/", relation.FriendList)                                   // 好友列表
	apiRouter.POST("/message/action/", message.MessageAction)                                      // 发送消息
	apiRouter.GET("/message/chat/", message.MessageChat)                                           // 聊天记录

}
