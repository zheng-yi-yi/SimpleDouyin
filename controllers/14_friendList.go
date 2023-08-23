package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// FriendList 好友列表
func FriendList(c *gin.Context) {
	// 获取传递的 token
	token := c.Query("token")

	// 使用 GetUserFromToken 函数获取已登录用户的信息
	user, exists := GetUserFromToken(token)
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{StatusCode: 1, StatusMsg: "未登录或登录已过期"})
		return
	}

	// 获取已登录用户的 ID
	userID := user.ID

	// 调用 socializeService.GetFriendsByUserID 通过用户ID获取好友列表
	friends, err := socializeService.GetFriendsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	// 构建返回的好友列表
	var userList UserList
	for _, friend := range friends {
		userList = append(userList, User{
			Id:            int64(friend.ID),
			Name:          friend.UserName,
			FollowCount:   int64(friend.FollowCount),
			FollowerCount: int64(friend.FollowerCount),
			IsFollow:      true, // 因为是好友关系，所以设置为 true
			Avatar:        config.AvatarURL,
			Background:    config.BackgroundURL,
			Signature:     config.SignatureStr,
		})
	}

	// 返回好友列表
	c.JSON(http.StatusOK, gin.H{
		"friends": userList,
	})
}

// GetUserFromToken 根据传递的 token 获取已登录用户的信息
func GetUserFromToken(token string) (*models.User, bool) {
	user, exists := UsersLoginInfo[token]
	if !exists {
		return nil, false
	}
	return &user, true
}
