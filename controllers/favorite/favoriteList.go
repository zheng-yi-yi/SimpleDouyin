package favorite

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/controllers/video"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// FavoriteList , 获取用户的所有点赞视频
func FavoriteList(c *gin.Context) {
	// 要查询的用户id
	query_user_id, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		// 用户id参数类型转换失败
		response.UserIdConversionError(c)
		return
	}
	// 根据要查询的用户id，取出该用户点赞的所有视频ID
	videoIds, err := FavoriteService.GetFavoriteList(uint(query_user_id))
	if err != nil {
		response.UnsuccessfulAction(c)
		return
	}
	// 根据点赞过的视频ID 取出所有对应的视频信息
	videoInfoList := video.VideoService.GetVideoInfoByIds(videoIds)
	// 如果点赞列表长度为0，则返回空的成功响应
	if len(videoInfoList) == 0 {
		// 视频列表获取成功
		response.GetFavoriteListSucceeded(c, []response.Video{})
		return
	}
	videoList := make([]response.Video, 0, len(videoInfoList))
	for i := 0; i < len(videoInfoList); i++ {
		videoList = append(videoList, response.Video{
			Id: int64(videoInfoList[i].ID),
			Author: response.User{
				Id:             int64(videoInfoList[i].User.ID),
				Name:           videoInfoList[i].User.UserName,
				FollowCount:    int64(videoInfoList[i].User.FollowCount),
				FollowerCount:  int64(videoInfoList[i].User.FollowerCount),
				IsFollow:       models.IsFollow(uint(query_user_id), videoInfoList[i].User.ID),
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
			IsFavorite:    models.IsFavorite(uint(query_user_id), videoInfoList[i].ID),
			Title:         videoInfoList[i].Description,
		})
	}
	// 获取视频列表成功
	response.GetFavoriteListSucceeded(c, videoList)
}
