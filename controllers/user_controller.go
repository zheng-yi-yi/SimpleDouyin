package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	Background    string `json:"background_image,omitempty"`
	Signature     string `json:"signature,omitempty"`
}

type UserList []User

var UsersLoginInfo = map[string]models.User{}

type UserLoginResponse struct {
	Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User `json:"user"`
}

// 处理用户的注册请求
func Register(c *gin.Context) {
	// 获取传递的用户名和密码。
	username := c.Query("username")
	password := c.Query("password")
	// 简单地将用户名和密码连接起来作为 token 的值，后续再改为更强的加密手段来生成 token
	token := username + password
	// 调用 userService.Register 来进行用户注册，并获取注册结果。
	user, err := userService.Register(username, password)
	// 检查注册结果，如果发生错误则返回错误信息：
	if err != nil {
		c.JSON(
			http.StatusOK,
			UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: err.Error()},
			})
		return
	}
	// 存储已注册的用户的信息
	UsersLoginInfo[token] = user
	// 返回成功的 JSON 响应给客户端，包含注册成功的状态信息、用户ID和 token
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		},
		UserId: int(user.ID),
		Token:  token,
	})
}

// 用户登录
func Login(c *gin.Context) {
	// 获取传递的用户名和密码。
	username := c.Query("username")
	password := c.Query("password")
	// 简单地将用户名和密码连接起来作为 token 的值，后续再改为更强的加密手段来生成 token
	token := username + password
	// 调用 userService.Login 来进行用户登录，并获取登录结果。
	user, err := userService.Login(username, password)
	// 检查登录结果，如果发生错误则返回错误信息：
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	// 存储已登录的用户的信息
	UsersLoginInfo[token] = user
	// 返回成功的 JSON 响应给客户端
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "登陆成功",
		},
		UserId: int(user.ID),
		Token:  token,
	})
}

// 获取用户信息
func UserInfo(c *gin.Context) {
	// 从上下文中获取已经鉴权的用户ID，并从查询参数中获取user_id，它是要查询的用户的ID。
	loginUserId, queryUserIdStr := c.GetUint("userID"), c.Query("user_id")
	// 解析并验证查询参数中的user_id，将其转换为uint64类型的值
	var queryUserId uint64
	if queryUserIdStr != "" {
		if _queryUserId, parseUintErr := strconv.ParseUint(queryUserIdStr, 10, 64); parseUintErr != nil {
			Failed(c, parseUintErr.Error())
			return
		} else {
			queryUserId = _queryUserId
		}
	}
	// 如果登录用户的ID为0，表示用户不存在或鉴权失败，会返回用户不存在的错误响应。
	if loginUserId == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "不存在该用户"})
		return
	}
	// 根据 queryUserId 的值，决定要查询哪个用户的信息。
	isFollow := false
	userId := loginUserId
	if queryUserId != 0 {
		userId = uint(queryUserId)
		isFollow = IsFollow(loginUserId, uint(queryUserId))
	}
	// 获取指定用户的信息
	userInfo, getUserInfoErr := userService.GetUserInfoById(userId)
	if getUserInfoErr != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: getUserInfoErr.Error()})
		return
	}
	// 构建用户信息的响应
	// 基本的响应，后续要更改，返回更加丰富的内容
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User: User{
			Id:            int64(userInfo.ID),
			Name:          userInfo.UserName,
			FollowCount:   int64(userInfo.FollowCount),
			FollowerCount: int64(userInfo.FollowerCount),
			IsFollow:      isFollow,
			Avatar:        config.AvatarURL,
			Background:    config.BackgroundURL,
			Signature:     config.SignatureStr,
		},
	})
}
