package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 发送消息响应
type MessageActionResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// MessageActionSucceeded , 发送消息成功
func MessageActionSucceeded(c *gin.Context) {
	c.JSON(http.StatusOK, MessageActionResponse{
		StatusCode: 0,
		StatusMsg:  "发送消息成功",
	})
}

// MessageActionError ， 发送消息失败
func MessageActionError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, MessageActionResponse{
		StatusCode: 1,
		StatusMsg:  "发送消息失败",
	})
}

type Message struct {
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID int64  `json:"from_user_id"` // 消息发送者id
	ID         int64  `json:"id"`           // 消息id
	ToUserID   int64  `json:"to_user_id"`   // 消息接收者id
}

// 聊天记录响应
type MessageChatResponse struct {
	StatusCode  string    `json:"status_code"`            // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg"`             // 返回状态描述
	MessageList []Message `json:"message_list,omitempty"` // 用户列表

}

// GetMessageChatSucceeded , 聊天记录获取成功
func GetMessageChatSucceeded(c *gin.Context, messageList []Message) {
	c.JSON(http.StatusOK, MessageChatResponse{
		StatusCode:  "0",
		StatusMsg:   "聊天记录获取成功",
		MessageList: messageList,
	})
}

// GetMessageChatError , 聊天记录获取失败
func GetMessageChatError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, MessageChatResponse{
		StatusCode: "0",
		StatusMsg:  "聊天记录获取失败",
	})
}
