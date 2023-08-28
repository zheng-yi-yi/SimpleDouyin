package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type User struct {
	Id             int64  `json:"id,omitempty"`               // 用户id
	Name           string `json:"name,omitempty"`             // 用户名称
	FollowCount    int64  `json:"follow_count,omitempty"`     // 关注总数
	FollowerCount  int64  `json:"follower_count,omitempty"`   // 粉丝总数
	IsFollow       bool   `json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Avatar         string `json:"avatar,omitempty"`           // 用户头像
	Background     string `json:"background_image,omitempty"` // 用户个人页顶部大图
	Signature      string `json:"signature,omitempty"`        // 个人简介
	FavoriteCount  int64  `json:"favorite_count"`             // 喜欢数
	TotalFavorited string `json:"total_favorited"`            // 获赞数量
	WorkCount      int64  `json:"work_count"`                 // 作品数
}

// 用户注册响应
type UserRegisterResponse struct {
	StatusCode int32  `json:"status_code"`       // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`        // 返回状态描述
	UserId     int32  `json:"user_id,omitempty"` // 用户鉴权token
	Token      string `json:"token,omitempty"`   // 用户id
}

// RegisterUserFailure ， 用户注册失败
func RegisterUserFailure(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, UserRegisterResponse{
		StatusCode: 1,
		StatusMsg:  "用户注册失败",
	})
}

// RegisterTokenError ， 用户注册时token生成失败
func RegisterTokenError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, UserRegisterResponse{
		StatusCode: 1,
		StatusMsg:  "token生成失败",
	})
}

// RegisterUserComplete ， 用户注册成功
func RegisterUserComplete(c *gin.Context, userId int32, token string) {
	c.JSON(http.StatusOK, UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "用户注册成功",
		UserId:     userId,
		Token:      token,
	})
}

// 用户登录响应
type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`       // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`        // 返回状态描述
	UserId     int32  `json:"user_id,omitempty"` // 用户鉴权token
	Token      string `json:"token,omitempty"`   // 用户id
}

// UserLoginFailure ， 用户登录失败
func UserLoginFailure(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, UserLoginResponse{
		StatusCode: 1,
		StatusMsg:  "用户登录失败",
	})
}

// UserLoginTokenError ， 用户登录时token生成失败
func UserLoginTokenError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, UserLoginResponse{
		StatusCode: 1,
		StatusMsg:  "token生成失败",
	})
}

// UserLoginComplete ， 用户登录成功
func UserLoginComplete(c *gin.Context, userId int32, token string) {
	c.JSON(http.StatusOK, UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "用户登录成功",
		UserId:     userId,
		Token:      token,
	})
}

// 用户信息响应
type UserInfoResponse struct {
	StatusCode int32  `json:"status_code"`    // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`     // 返回状态描述
	User       User   `json:"user,omitempty"` // 用户信息
}

// GetUserInfoFailure ， 用户信息获取失败
func GetUserInfoFailure(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, UserInfoResponse{
		StatusCode: 1,
		StatusMsg:  "用户信息获取失败",
	})
}

// UserInfoConversionError ， 参数类型转换失败
func UserInfoConversionError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, UserInfoResponse{
		StatusCode: 1,
		StatusMsg:  "参数类型转换失败",
	})
}

// UserInfoComplete ， 用户信息获取成功
func UserInfoComplete(c *gin.Context, userInfo models.User, userId uint, query_user_id uint64) {
	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "用户信息获取成功",
		User: User{
			Id:             int64(userInfo.ID),
			Name:           userInfo.UserName,
			FollowCount:    int64(userInfo.FollowCount),
			FollowerCount:  int64(userInfo.FollowerCount),
			IsFollow:       models.IsFollow(userId, uint(query_user_id)),
			Avatar:         userInfo.Avatar,
			Background:     userInfo.BackgroundImage,
			Signature:      userInfo.Signature,
			TotalFavorited: userInfo.TotalFavorited,
			WorkCount:      userInfo.WorkCount,
			FavoriteCount:  userInfo.FavoriteCount,
		},
	})
}
