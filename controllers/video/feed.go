package video

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/controllers/favorite"
	"github.com/zheng-yi-yi/simpledouyin/controllers/relation"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/services"
	"github.com/zheng-yi-yi/simpledouyin/utils"
)

var VideoService services.VideoService

// 视频流
func Feed(c *gin.Context) {
	token := c.Query("token")
	lastTimestamp := c.Query("latest_time")
	startTime := utils.CalculateStartTime(lastTimestamp)
	userId := response.UsersLoginInfo[token].ID
	feedVideo := *VideoService.Feed(startTime)
	lenFeedVideoList := len(feedVideo)
	if lenFeedVideoList == 0 {
		// 空数据处理
		c.JSON(http.StatusOK, response.FeedResponse{
			Response:  response.Response{StatusCode: 0},
			VideoList: []response.Video{},
			NextTime:  time.Now().Unix(),
		})
		return
	}
	nextTime := feedVideo[lenFeedVideoList-1].CreatedAt.Unix()
	videoList := make([]response.Video, 0, lenFeedVideoList)
	for _, video := range feedVideo {
		videoList = append(videoList,
			response.Video{
				Id: int64(video.ID),
				Author: response.User{
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
				PlayUrl:       config.SERVER_RESOURCES + video.PlayUrl,
				CoverUrl:      config.SERVER_RESOURCES + video.CoverUrl,
				FavoriteCount: video.FavoriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    IsFavorite(userId, video.ID),
				Title:         video.Description,
			})
	}
	c.JSON(http.StatusOK, response.FeedResponse{
		Response:  response.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}

// 判断登录用户是否关注了视频作者
func IsFollow(fromUserId, toUserId uint) bool {
	if fromUserId == 0 {
		return false
	}
	return relation.RelationService.IsFollow(fromUserId, toUserId)
}

// 判断用户是否点赞当前视频
func IsFavorite(userId, videoId uint) bool {
	if userId == 0 {
		return false
	}
	return favorite.FavoriteService.IsFavorite(userId, videoId)
}
