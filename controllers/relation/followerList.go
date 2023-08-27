package relation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/services"
)

var RelationService services.RelationService

// FollowerList 拉取粉丝用户列表
func FollowerList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	users, err := RelationService.GetFollowerList(uint(userId))
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	var relaUserList []response.RelationUser
	for _, user := range users {
		isFollow := RelationService.IsFollow(uint(userId), user.ID)
		RelationUser := response.RelationUser{
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
		relaUserList = append(relaUserList, RelationUser)
	}
	c.JSON(http.StatusOK, response.UserListResponse{
		Response: response.Response{StatusCode: 0},
		UserList: relaUserList,
	})
}
