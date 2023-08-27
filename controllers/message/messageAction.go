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
	// 当前用户id
	token := c.Query("token")
	from_user_id := response.UsersLoginInfo[token].ID
	// 对方用户id
	to_user_id, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	action_type := c.Query("action_type")
	// 发送信息
	if action_type == "1" {
		content := c.Query("content") // 消息内容
		if err := MessageService.AddMessage(uint(from_user_id), content, uint(to_user_id)); err != nil {
			response.Failed(c, err.Error())
			return
		}
		response.Success(c, "成功发送信息")
		return
	}
	response.Failed(c, "消息发送失败")
}
