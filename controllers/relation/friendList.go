package relation

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// FriendList ， 获取好友列表
func FriendList(c *gin.Context) {
	// 当前登录的用户
	userId := c.Value("userID").(uint)
	// 要查询的用户id
	query_user_id, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		response.UserIdConversionError(c) // 用户id参数类型转换失败
		return
	}
	// 获取好友列表
	users, err := RelationService.GetFriendsList(uint(query_user_id))
	if err != nil {
		response.GetFriendListError(c) // 好友列表获取失败
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
	response.GetFriendListSucceeded(c, relationUsers) // 好友列表获取成功
}
