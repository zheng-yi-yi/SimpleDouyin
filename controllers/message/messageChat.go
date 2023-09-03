package message

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
)

// MessageChat ，获取聊天记录
func MessageChat(c *gin.Context) {
	// 当前登录的用户
	from_user_id := c.Value("userID").(uint)
	// 获取对方用户id
	to_user_id, err := strconv.ParseUint(c.Query("to_user_id"), 10, 64)
	if err != nil {
		response.ToUserIdConversionError(c) // 对方用户id参数类型转换失败
		return
	}
	// 获取pre_msg_time
	preMsgTimeUnix, err := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
	if err != nil {
		response.PreMsgTimeConversionError(c)
		return
	}
	preMsgTime := time.Unix(0, preMsgTimeUnix*int64(time.Millisecond))
	// 获取消息列表
	messages, err := MessageService.GetMessageListWithTime(from_user_id, uint(to_user_id), preMsgTime)
	if err != nil {
		response.GetMessageChatError(c) // 聊天记录获取失败
		return
	}
	messageList := make([]response.Message, 0, len(messages))
	for _, message := range messages {
		messageList = append(messageList, response.Message{
			ID:         int64(message.ID),
			ToUserID:   int64(message.ToUserID),
			FromUserID: int64(message.FromUserID),
			Content:    message.Content,
			CreateTime: message.CreateTime.Unix(),
		})
	}
	response.GetMessageChatSucceeded(c, messageList) // 聊天记录获取成功
}
