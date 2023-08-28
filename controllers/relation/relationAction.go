package relation

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
)

// RelationAction ， 关注操作
func RelationAction(c *gin.Context) {
	// 当前登录的用户
	formUserId := c.Value("userID").(uint)
	// 要查询的对方用户id
	toUserId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 64)
	if err != nil {
		response.ToUserIdConversionError(c)
		return
	}
	actionType := c.Query("action_type")
	if toUserId == uint64(formUserId) {
		response.UnsuccessfulAction(c) // 操作失败（自己不用关注或取关自己）
		return
	}
	switch actionType {
	case "1":
		//关注操作
		err := RelationService.FollowUser(uint(formUserId), uint(toUserId))
		if err != nil {
			response.RelationActionError(c) // 关注失败
		}
		response.RelationActionSucceeded(c) // 关注成功
		return
	case "2":
		//取消关注操作
		err := RelationService.CancelFollowUser(uint(formUserId), uint(toUserId))
		if err != nil {
			response.CancelRelationError(c) // 取关失败
		}
		response.CancelRelationSucceeded(c) // 取关成功
		return
	default:
		response.UnsuccessfulAction(c) // 操作失败
	}
}
