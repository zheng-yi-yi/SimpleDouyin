package video

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/middlewares"
	"github.com/zheng-yi-yi/simpledouyin/models"
	"github.com/zheng-yi-yi/simpledouyin/services"
	"github.com/zheng-yi-yi/simpledouyin/utils"
)

var VideoService services.VideoService

// Feed , 获取视频流
func Feed(c *gin.Context) {
	tokenString := c.Query("token")
	if len(tokenString) == 0 {
		// 未登录状态下的视频流获取
		NoLoginAccess(c)
		return
	}
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &middlewares.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return middlewares.JwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, middlewares.AuthFailResponse{StatusCode: 1, StatusMsg: middlewares.ParseTokenFailed})
		c.Abort()
		return
	}
	// 校验鉴权的声明
	claims, ok := token.Claims.(*middlewares.MyClaims)
	if ok && token.Valid {
		// 已登录状态下的视频流获取
		LoginAccess(c, claims.UserID)
		return
	}
	c.JSON(http.StatusInternalServerError, middlewares.AuthFailResponse{StatusCode: 1, StatusMsg: middlewares.CheckFailed})
}

// 未登录时的视频流获取
func NoLoginAccess(c *gin.Context) {
	startTime := utils.CalculateStartTime(c.Query("latest_time"))
	feedVideo := *VideoService.Feed(startTime)
	lenFeedVideoList := len(feedVideo)
	// 空数据处理
	if lenFeedVideoList == 0 {
		response.GetFeedSuccess(c, time.Now().Unix(), []response.Video{})
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
				Title:         video.Description,
			})
	}
	response.GetFeedSuccess(c, nextTime, videoList)
}

// 未登录时的视频流获取
func LoginAccess(c *gin.Context, userId uint) {
	startTime := utils.CalculateStartTime(c.Query("latest_time"))
	feedVideo := *VideoService.Feed(startTime)
	lenFeedVideoList := len(feedVideo)
	if lenFeedVideoList == 0 {
		// 空数据处理
		response.GetFeedSuccess(c, time.Now().Unix(), []response.Video{})
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
					IsFollow:       models.IsFollow(userId, video.User.ID),
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
				IsFavorite:    models.IsFavorite(userId, video.ID),
				Title:         video.Description,
			})
	}
	response.GetFeedSuccess(c, nextTime, videoList)
}
