package relation

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// FollowList ，关注列表
func FollowList(c *gin.Context) {
	// 当前登录的用户
	userId := c.Value("userID").(uint)
	// 要查询的用户id
	query_user_id, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		response.UserIdConversionError(c) // 用户id参数类型转换失败
		return
	}
	// 查询用户关注的所有用户
	users, err := RelationService.GetFllowList(uint(query_user_id))
	if err != nil {
		// 关注列表获取失败
		response.GetFollowListError(c)
		return
	}
	var relationUsers []response.User
	for _, user := range users {
		relationUser := response.User{
			Id:             int64(user.ID),
			Name:           user.UserName,
			Avatar:         user.Avatar,
			Signature:      user.Signature,
			FollowCount:    int64(user.FollowCount),
			FollowerCount:  int64(user.FollowCount),
			IsFollow:       models.IsFollow(uint(userId), user.ID),
			Background:     user.BackgroundImage,
			TotalFavorited: user.TotalFavorited,
			WorkCount:      user.WorkCount,
			FavoriteCount:  user.FavoriteCount,
		}
		relationUsers = append(relationUsers, relationUser)
	}
	// 关注列表获取成功
	response.GetFollowListSucceeded(c, relationUsers)
}
