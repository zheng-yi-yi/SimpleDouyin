package favorite

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/services"
)

var FavoriteService services.FavoriteService

// FavoriteAction ，登录用户对视频的点赞和取消点赞操作
func FavoriteAction(c *gin.Context) {
	video_id, err := strconv.ParseUint(c.Query("video_id"), 10, 64)
	if err != nil {
		// 视频id参数类型转换失败
		response.VideoIdConversionError(c)
		return
	}
	userId := c.Value("userID").(uint)
	actionType := c.Query("action_type")
	// 判断操作是否合法
	switch actionType {
	case "1":
		// 点赞操作
		err := FavoriteService.AddLike(userId, uint(video_id))
		if err != nil {
			// 点赞失败
			response.LikeActionError(c)
			return
		}
		// 点赞成功
		response.LikeActionSucceeded(c)
		return
	case "2":
		// 取消赞操作
		err := FavoriteService.CancelLike(userId, uint(video_id))
		if err != nil {
			// 取消赞失败
			response.CancelLikeError(c)
			return
		}
		// 取消赞成功
		response.UnlikeActionSucceeded(c)
		return
	default:
		// 操作不合法
		response.UnsuccessfulAction(c)
	}
}
