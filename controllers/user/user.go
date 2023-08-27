package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/relation"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
)

// 获取用户信息
func UserInfo(c *gin.Context) {
	// 从上下文中获取已经鉴权的用户ID，并从查询参数中获取user_id，它是要查询的用户的ID。
	loginUserId, queryUserIdStr := c.GetUint("userID"), c.Query("user_id")
	// 解析并验证查询参数中的user_id，将其转换为uint64类型的值
	var queryUserId uint64
	if queryUserIdStr != "" {
		if _queryUserId, parseUintErr := strconv.ParseUint(queryUserIdStr, 10, 64); parseUintErr != nil {
			response.Failed(c, parseUintErr.Error())
			return
		} else {
			queryUserId = _queryUserId
		}
	}
	// 如果登录用户的ID为0，表示用户不存在或鉴权失败，会返回用户不存在的错误响应。
	if loginUserId == 0 {
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "不存在该用户"})
		return
	}
	// 根据 queryUserId 的值，决定要查询哪个用户的信息。
	isFollow := false
	userId := loginUserId
	if queryUserId != 0 {
		userId = uint(queryUserId)
		isFollow = relation.RelationService.IsFollow(loginUserId, uint(queryUserId))
	}
	// 获取指定用户的信息
	userInfo, getUserInfoErr := UserService.GetUserInfoById(userId)
	if getUserInfoErr != nil {
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: getUserInfoErr.Error()})
		return
	}
	// 构建用户信息的响应
	c.JSON(http.StatusOK, response.UserResponse{
		Response: response.Response{StatusCode: 0},
		User: response.User{
			Id:             int64(userInfo.ID),
			Name:           userInfo.UserName,
			FollowCount:    int64(userInfo.FollowCount),
			FollowerCount:  int64(userInfo.FollowerCount),
			IsFollow:       isFollow,
			Avatar:         userInfo.Avatar,
			Background:     userInfo.BackgroundImage,
			Signature:      userInfo.Signature,
			TotalFavorited: userInfo.TotalFavorited,
			WorkCount:      userInfo.WorkCount,
			FavoriteCount:  userInfo.FavoriteCount,
		},
	})
}
