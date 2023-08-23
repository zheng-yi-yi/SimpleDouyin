package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
