package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/middlewares"
)

// Register , 处理用户的注册请求
func Register(c *gin.Context) {
	// 获取请求参数
	username := c.Query("username")
	password := c.Query("password")
	// 用户注册
	user, err := UserService.Register(username, password)
	if err != nil {
		// 用户注册失败
		response.RegisterUserFailure(c)
		return
	}
	// 生成 token
	token, err := middlewares.GenerateToken(user.ID, username, password)
	if err != nil {
		// token 生成失败
		response.RegisterTokenError(c)
		return
	}
	// 用户注册成功
	response.RegisterUserComplete(c, int32(user.ID), token)
}
