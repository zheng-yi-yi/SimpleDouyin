package video

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// 获取用户发布视频列表的处理函数
func PublishList(c *gin.Context) {
	// 当前用户
	userId := c.Value("userID").(uint)
	// 要查询的用户id
	query_user_id, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		// 用户id参数类型转换失败
		response.UserIdConversionError(c)
		return
	}
	// 获取用户发布的视频列表。
	userPublishList := VideoService.UserPublishList(uint(query_user_id))
	// 如果用户发布列表为空，返回一个空的成功响应。
	if len(userPublishList) == 0 {
		response.GetPublishListSuccess(c, []response.Video{})
		return
	}
	// 创建视频列表：遍历用户发布列表，将每个视频对象映射到一个新的视频对象，同时进行一些字段的转换和处理。
	videoList := make([]response.Video, 0, len(userPublishList))
	for i := 0; i < len(userPublishList); i++ {
		videoList = append(videoList, response.Video{
			Id: int64(userPublishList[i].ID),
			Author: response.User{
				Id:             int64(userPublishList[i].User.ID),
				Name:           userPublishList[i].User.UserName,
				FollowCount:    int64(userPublishList[i].User.FollowCount),
				FollowerCount:  int64(userPublishList[i].User.FollowerCount),
				IsFollow:       models.IsFollow(userId, uint(query_user_id)),
				Avatar:         userPublishList[i].User.Avatar,
				Background:     userPublishList[i].User.BackgroundImage,
				Signature:      userPublishList[i].User.Signature,
				TotalFavorited: userPublishList[i].User.TotalFavorited,
				WorkCount:      userPublishList[i].User.WorkCount,
				FavoriteCount:  userPublishList[i].User.FavoriteCount,
			},
			PlayUrl:       config.SERVER_RESOURCES + userPublishList[i].PlayUrl,
			CoverUrl:      config.SERVER_RESOURCES + userPublishList[i].CoverUrl,
			FavoriteCount: userPublishList[i].FavoriteCount,
			CommentCount:  userPublishList[i].CommentCount,
			IsFavorite:    models.IsFavorite(userId, uint(query_user_id)),
			Title:         userPublishList[i].Description,
		})
	}
	// 最后，返回带有视频列表的成功响应。
	response.GetPublishListSuccess(c, videoList)
}
