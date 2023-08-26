package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction ，登录用户对视频的点赞和取消点赞操作
func FavoriteAction(c *gin.Context) {
	// 获取当前用户
	userId := c.GetUint("userID")
	if userId == 0 {
		Failed(c, "用户不存在")
		return
	}
	actionType := c.Query("action_type")
	video_id, err := strconv.ParseUint(c.Query("video_id"), 10, 64)
	if err != nil {
		Failed(c, "获取到非法的视频ID")
		return
	}
	// 判断操作是否合法
	if actionType != "1" && actionType != "2" {
		Failed(c, "非法操作")
		return
	}
	// 点赞操作
	if actionType == "1" {
		err := favoriteService.AddLike(userId, uint(video_id))
		if err != nil {
			Failed(c, err.Error())
			return
		}
		Success(c, "操作成功")
		return
	}
	// 取消赞操作
	err = favoriteService.CancelLike(userId, uint(video_id))
	if err != nil {
		Failed(c, err.Error())
		return
	}
	Success(c, "操作成功")
}
