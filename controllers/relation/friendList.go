package relation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
)

// FriendList 获取好友列表
func FriendList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	users, err := RelationService.GetFriendsList(uint(userId))
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	var relationUsers []response.RelationUser
	for _, user := range users {
		isFollow := RelationService.IsFollow(uint(userId), user.ID)
		relationUser := response.RelationUser{
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
	c.JSON(http.StatusOK, response.UserListResponse{
		Response: response.Response{StatusCode: 0},
		UserList: relationUsers,
	})
}
