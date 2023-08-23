package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
