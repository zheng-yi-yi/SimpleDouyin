package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MessageChat ，获取聊天记录
func MessageChat(c *gin.Context) {
	// 获取当前用户id
	token := c.Query("token")
	userId := UsersLoginInfo[token].ID
	// 获取对方用户id
	to_user_id, err := strconv.ParseUint(c.Query("to_user_id"), 10, 64)
	if err != nil {
		Failed(c, err.Error())
		return
	}
	// 获取消息列表
	messages, err := messageService.GetMessageList(userId, uint(to_user_id))
	if err != nil {
		c.JSON(http.StatusOK, ChatResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "无法获取聊天记录",
			},
			MessageList: nil,
		})
		return
	}
	messageList := make([]Message, 0, len(messages))
	for _, message := range messages {
		messageList = append(messageList, Message{
			ID:         int64(message.ID),
			ToUserID:   int64(message.ToUserID),
			FromUserID: int64(message.FromUserID),
			Content:    message.Content,
			CreateTime: message.CreateTime.Unix(),
		})
	}
	c.JSON(http.StatusOK, ChatResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "成功获取聊天记录",
		},
		MessageList: messageList,
	})
}
