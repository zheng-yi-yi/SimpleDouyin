package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

// 关注列表
func FollowList(c *gin.Context) {
	userIdStr := c.Query("user_id")

	var formUserId uint64

	_formUserId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		Failed(c, err.Error())
		return
	} else {
		formUserId = _formUserId
	}

	//找用户关注的所有用户
	users, err := relationService.GetFllowList(uint(formUserId))
	if err != nil {
		Failed(c, err.Error())
		return
	}

	var relationUsers []relationUser
	//返回的用户信息到时候再完善一下TotalFavorited、WorkCount、FavoriteCount
	for _, user := range users {
		isFollow := relationService.IsFollow(uint(formUserId), user.ID)
		relationUser := relationUser{
			ID:              int64(user.ID),
			Name:            user.UserName,
			Avatar:          config.AvatarURL,
			Signature:       config.SignatureStr,
			FollowCount:     int64(user.FollowCount),
			FollowerCount:   int64(user.FollowerCount),
			IsFollow:        isFollow,
			BackgroundImage: config.BackgroundURL,
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
