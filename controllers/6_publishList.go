package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

// 获取用户发布视频列表的处理函数
func PublishList(c *gin.Context) {
	// 从上下文中获取已经鉴权的用户ID。
	userId := c.GetUint("userID")
	// 如果用户ID为0，表示用户不存在或鉴权失败，会返回相应的错误响应。
	if userId == 0 {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "不存在该用户",
			},
			VideoList: []Video{},
		})
		return
	}
	// 调用 service 层中的方法，获取用户发布的视频列表。
	userPublishList := videoService.UserPublishList(userId)
	// 如果用户发布列表为空，返回一个空的成功响应。
	if len(userPublishList) == 0 {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			VideoList: []Video{},
		})
		return
	}
	// 创建视频列表：遍历用户发布列表，将每个视频对象映射到一个新的视频对象，同时进行一些字段的转换和处理。
	videoList := make([]Video, 0, len(userPublishList))
	for i := 0; i < len(userPublishList); i++ {
		videoList = append(videoList, Video{
			Id: int64(userPublishList[i].ID),
			Author: User{
				Id:            int64(userPublishList[i].User.ID),
				Name:          userPublishList[i].User.UserName,
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			},
			PlayUrl:       config.ResourceServerURL + userPublishList[i].PlayUrl,
			CoverUrl:      config.ResourceServerURL + userPublishList[i].CoverUrl,
			FavoriteCount: userPublishList[i].FavoriteCount,
			CommentCount:  userPublishList[i].CommentCount,
			IsFavorite:    false,
			Title:         userPublishList[i].Description,
		})
	}
	// 最后，返回带有视频列表的成功响应。
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videoList,
	})
}
