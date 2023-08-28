package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// 获取用户信息
func UserInfo(c *gin.Context) {
	// 当前登录的用户
	userId := c.Value("userID").(uint)
	// 要获取用户信息的用户id
	query_user_id, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		// 参数类型转换失败
		response.UserInfoConversionError(c)
		return
	}
	// 获取指定用户的信息
	userInfo, getUserInfoErr := models.FetchData(uint(query_user_id))
	if getUserInfoErr != nil {
		// 用户信息获取失败
		response.GetUserInfoFailure(c)
		return
	}
	// 构建用户信息的响应
	response.UserInfoComplete(c, userInfo, userId, query_user_id)
}
