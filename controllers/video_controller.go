package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/utils"
)

// 视频流
func Feed(c *gin.Context) {
	token := c.Query("token")
	lastTimestamp := c.Query("latest_time")
	startTime := utils.CalculateStartTime(lastTimestamp)
	userId := UsersLoginInfo[token].ID
	feedVideo := *videoService.Feed(startTime)
	lenFeedVideoList := len(feedVideo)
	if lenFeedVideoList == 0 {
		// 空数据处理
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: []Video{},
			NextTime:  time.Now().Unix(),
		})
		return
	}
	nextTime := feedVideo[lenFeedVideoList-1].CreatedAt.Unix()
	videoList := make([]Video, 0, lenFeedVideoList)
	for _, video := range feedVideo {
		videoList = append(videoList,
			Video{
				Id: int64(video.ID),
				Author: User{
					Id:            int64(video.User.ID),
					Name:          video.User.UserName,
					FollowCount:   int64(video.User.FollowerCount),
					FollowerCount: int64(video.User.FollowerCount),
					IsFollow:      IsFollow(userId, video.User.ID),
					Avatar:        config.AvatarURL,
					Background:    config.BackgroundURL,
					Signature:     config.SignatureStr,
				},
				PlayUrl:       config.ResourceServerURL + video.PlayUrl,
				CoverUrl:      config.ResourceServerURL + video.CoverUrl,
				FavoriteCount: video.FavoriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    IsFavorite(userId, video.ID),
				Title:         video.Description,
			})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}

// 视频投稿发布
func Publish(c *gin.Context) {

}

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

// ================= 功能函数 =================

// 判断用户是否点赞当前视频
func IsFavorite(userId, videoId uint) bool {
	if userId == 0 {
		return false
	}
	return favoriteService.IsFavorite(userId, videoId)
}

// 判断登录用户是否关注了视频作者
func IsFollow(fromUserId, toUserId uint) bool {
	if fromUserId == 0 {
		return false
	}
	return relationService.IsFollow(fromUserId, toUserId)
}
