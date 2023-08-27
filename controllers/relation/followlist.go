package relation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
)

// 关注列表
func FollowList(c *gin.Context) {
	loginUserId := c.GetUint("userID")
	if loginUserId == 0 {
		response.Failed(c, "用户不存在")
		return
	}
	userIdStr := c.Query("user_id")
	var formUserId uint64
	_formUserId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		response.Failed(c, err.Error())
		return
	} else {
		formUserId = _formUserId
	}
	// 查询用户关注的所有用户
	users, err := RelationService.GetFllowList(uint(formUserId))
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	var relationUsers []response.RelationUser
	for _, user := range users {
		isFollow := RelationService.IsFollow(uint(formUserId), user.ID)
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
