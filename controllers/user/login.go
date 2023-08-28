package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/middlewares"
	"github.com/zheng-yi-yi/simpledouyin/services"
)

var UserService services.UserService

// Login , 用户登录
func Login(c *gin.Context) {
	// 获取请求参数
	username := c.Query("username")
	password := c.Query("password")
	// 用户登录
	user, err := UserService.Login(username, password)
	if err != nil {
		// 用户登录失败
		response.UserLoginFailure(c)
		return
	}
	// 生成 token
	token, err := middlewares.GenerateToken(user.ID, username, password)
	if err != nil {
		// 用户登录时 token 生成失败
		response.UserLoginTokenError(c)
		return
	}
	// 用户登录成功
	response.UserLoginComplete(c, int32(user.ID), token)
}
