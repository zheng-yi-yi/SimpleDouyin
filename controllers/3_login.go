package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
