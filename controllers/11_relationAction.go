package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 关注操作
func RelationAction(c *gin.Context) {

	toUserIdStr := c.Query("to_user_id")
	actionType := c.Query("action_type")
	//userId := c.GetUint("UserID")
	token := c.Query("token")
	userId := UsersLoginInfo[token].ID

	if userId == 0 {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "不存在该用户",
			},
			UserList: []relationUser{},
		})
		return
	}

	//获取存储到上下文的用户id
	formUserId := userId

	//获取请求参数中的被关注用户id

	var toUserId uint64

	_toUserId, err := strconv.ParseUint(toUserIdStr, 10, 64)
	if err != nil {
		Failed(c, err.Error())
		return
	} else {
		toUserId = _toUserId
	}

	switch actionType {
	case "1":
		//关注操作
		err := relationService.FollowUser(uint(formUserId), uint(toUserId))
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 3, StatusMsg: "关注失败"})
		}

	case "2":
		//取消关注操作
		err := relationService.CancelFollowUser(uint(formUserId), uint(toUserId))
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 3, StatusMsg: err.Error()})

		}
	default:
		c.JSON(http.StatusBadRequest, Response{StatusCode: 3, StatusMsg: "无效操作"})
	}
}
