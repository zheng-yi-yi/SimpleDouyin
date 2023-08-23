package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
