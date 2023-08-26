package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

// FavoriteList , 获取用户的所有点赞视频
func FavoriteList(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		Failed(c, err.Error())
		return
	}
	// 根据用户ID 取出该用户点赞的所有视频ID
	videoIds, err := favoriteService.GetFavoriteList(uint(userId))
	if err != nil {
		Failed(c, err.Error())
		return
	}
	// 根据点赞过的视频ID 取出所有对应的视频信息
	videoInfoList := videoService.GetVideoInfoByIds(videoIds)
	if len(videoInfoList) == 0 {
		c.JSON(http.StatusOK, ResponseVideoList{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			VideoList: []Video{},
		})
		return
	}
	videoList := make([]Video, 0, len(videoInfoList))
	for i := 0; i < len(videoInfoList); i++ {
		videoList = append(videoList, Video{
			Id: int64(videoInfoList[i].ID),
			Author: User{
				Id:             int64(videoInfoList[i].User.ID),
				Name:           videoInfoList[i].User.UserName,
				FollowCount:    int64(videoInfoList[i].User.FollowCount),
				FollowerCount:  int64(videoInfoList[i].User.FollowerCount),
				IsFollow:       IsFollow(uint(userId), videoInfoList[i].User.ID),
				Avatar:         videoInfoList[i].User.Avatar,
				Background:     videoInfoList[i].User.BackgroundImage,
				Signature:      videoInfoList[i].User.Signature,
				TotalFavorited: videoInfoList[i].User.TotalFavorited,
				WorkCount:      videoInfoList[i].User.WorkCount,
				FavoriteCount:  videoInfoList[i].User.FavoriteCount,
			},
			PlayUrl:       config.SERVER_RESOURCES + videoInfoList[i].PlayUrl,
			CoverUrl:      config.SERVER_RESOURCES + videoInfoList[i].CoverUrl,
			FavoriteCount: videoInfoList[i].FavoriteCount,
			CommentCount:  videoInfoList[i].CommentCount,
			IsFavorite:    IsFavorite(uint(userId), videoInfoList[i].ID),
			Title:         videoInfoList[i].Description,
		})
	}
	c.JSON(http.StatusOK, ResponseVideoList{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videoList,
	})
}
