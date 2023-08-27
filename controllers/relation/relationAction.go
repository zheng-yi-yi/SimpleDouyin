package relation

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
)

// 关注操作
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	userId := response.UsersLoginInfo[token].ID
	if userId == 0 {
		response.Failed(c, "用户不存在")
		return
	}
	toUserIdStr := c.Query("to_user_id")
	actionType := c.Query("action_type")
	//获取请求参数中的被关注用户id
	toUserId, parseUintErr := strconv.ParseUint(toUserIdStr, 10, 64)
	if parseUintErr != nil {
		response.Failed(c, parseUintErr.Error())
		return
	}
	//获取存储到上下文的用户id
	formUserId := userId
	// 自己不能关注/取消关注自己
	if toUserId == uint64(formUserId) {
		response.Failed(c, "无法关注自己")
		return
	}
	switch actionType {
	case "1":
		//关注操作
		err := RelationService.FollowUser(uint(formUserId), uint(toUserId))
		if err != nil {
			response.Failed(c, "关注失败")
		}
	case "2":
		//取消关注操作
		err := RelationService.CancelFollowUser(uint(formUserId), uint(toUserId))
		if err != nil {
			response.Failed(c, "取关失败")
		}
	default:
		response.Failed(c, "无效操作")
	}
	response.Success(c, "操作成功")
}
