package message

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/services"
)

var MessageService services.MessageService

// MessageAction ，发送消息
func MessageAction(c *gin.Context) {
	// 当前登录的用户
	from_user_id := c.Value("userID").(uint)
	// 要查询的对方用户id
	to_user_id, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		response.UserIdConversionError(c) // 用户id参数类型转换失败
		return
	}
	action_type := c.Query("action_type")
	// 发送信息
	if action_type == "1" {
		content := c.Query("content") // 消息内容
		if err := MessageService.AddMessage(uint(from_user_id), content, uint(to_user_id)); err != nil {
			response.MessageActionError(c) // 消息发送失败
			return
		}
		response.MessageActionSucceeded(c) // 消息发送成功
		return
	}
	response.UnsuccessfulAction(c) // 非法操作
}
