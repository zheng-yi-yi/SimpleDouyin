package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type ChatResponse struct {
	Response
	MessageList []models.Message `json:"message_list"`
}

// 发送消息
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	content := c.Query("content")

	if actionType != "1" {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 0, StatusMsg: "action_type不合法"})
	}
	if user, exist := UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		err := messageService.AddMessage(user.ID, uint(userIdB), content)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "消息发送失败"})
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "用户不存在"})
	}
}

// 消息记录
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	// preMsgTime := c.Query("pre_msg_time")

	if user, exist := UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		messageList, _ := messageService.MessageList(user.ID, uint(userIdB))
		c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: messageList})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
