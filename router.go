package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers"
	"github.com/zheng-yi-yi/simpledouyin/middlewares"
)

// 路由初始化函数
func initRouter(r *gin.Engine) {
	// 基本设置
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	apiRouter := r.Group("/douyin")
	// （1）基础功能：
	apiRouter.GET("/feed/", controllers.Feed)                                              //  视频流
	apiRouter.POST("/user/register/", controllers.Register)                                // 用户注册
	apiRouter.POST("/user/login/", controllers.Login)                                      // 用户登录
	apiRouter.GET("/user/", middlewares.JWTAuth("query"), controllers.UserInfo)            // 用户信息
	apiRouter.POST("/publish/action/", middlewares.JWTAuth("form"), controllers.Publish)   // 视频投稿
	apiRouter.GET("/publish/list/", middlewares.JWTAuth("query"), controllers.PublishList) // 发布列表
	// （2）互动功能：
	apiRouter.POST("/favorite/action/", middlewares.JWTAuth("query"), controllers.FavoriteAction) // 点赞操作
	apiRouter.GET("/favorite/list/", middlewares.JWTAuth("query"), controllers.FavoriteList)      // 喜欢列表
	apiRouter.POST("/comment/action/", middlewares.JWTAuth("query"), controllers.CommentAction)   // 评论操作
	apiRouter.GET("/comment/list/", controllers.CommentList)                                      // 评论列表
	// （3）社交功能：
	apiRouter.POST("/relation/action/", middlewares.JWTAuth("query"), controllers.RelationAction)     // 关注操作
	apiRouter.GET("/relation/follow/list/", middlewares.JWTAuth("query"), controllers.FollowList)     // 关注列表
	apiRouter.GET("/relation/follower/list/", middlewares.JWTAuth("query"), controllers.FollowerList) // 粉丝列表
	apiRouter.GET("/relation/friend/list/", controllers.FriendList)                                   // 好友列表
	apiRouter.POST("/message/action/", controllers.MessageAction)                                     // 发送消息
	apiRouter.GET("/message/chat/", controllers.MessageChat)                                          // 聊天记录

}
