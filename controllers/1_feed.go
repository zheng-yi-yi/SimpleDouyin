package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
					Id:             int64(video.User.ID),
					Name:           video.User.UserName,
					FollowCount:    int64(video.User.FollowCount),
					FollowerCount:  int64(video.User.FollowerCount),
					IsFollow:       IsFollow(userId, video.User.ID),
					Avatar:         video.User.Avatar,
					Background:     video.User.BackgroundImage,
					Signature:      video.User.Signature,
					TotalFavorited: video.User.TotalFavorited,
					WorkCount:      video.User.WorkCount,
					FavoriteCount:  video.User.FavoriteCount,
				},
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
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

// 判断登录用户是否关注了视频作者
func IsFollow(fromUserId, toUserId uint) bool {
	if fromUserId == 0 {
		return false
	}
	return relationService.IsFollow(fromUserId, toUserId)
}

// 判断用户是否点赞当前视频
func IsFavorite(userId, videoId uint) bool {
	if userId == 0 {
		return false
	}
	return favoriteService.IsFavorite(userId, videoId)
}
