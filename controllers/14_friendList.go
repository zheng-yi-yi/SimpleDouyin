package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FriendList 获取好友列表
func FriendList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		Failed(c, err.Error())
		return
	}
	users, err := relationService.GetFriendsList(uint(userId))
	if err != nil {
		Failed(c, err.Error())
		return
	}
	var relationUsers []relationUser
	for _, user := range users {
		isFollow := relationService.IsFollow(uint(userId), user.ID)
		relationUser := relationUser{
			ID:              int64(user.ID),
			Name:            user.UserName,
			Avatar:          user.Avatar,
			Signature:       user.Signature,
			FollowCount:     int64(user.FollowCount),
			FollowerCount:   int64(user.FollowCount),
			IsFollow:        isFollow,
			BackgroundImage: user.BackgroundImage,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
		relationUsers = append(relationUsers, relationUser)
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{StatusCode: 0},
		UserList: relationUsers,
	})
}
