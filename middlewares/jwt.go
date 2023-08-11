package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers"
)

// 简单的 JWT 鉴权函数：用于指定从哪个位置获取token（query 或 form）
func JWTAuth(where string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		if where == "form" {
			token = c.PostForm("token")
		} else {
			token = c.Query("token")
		}
		// 不存在该用户token则直接抛出用户不存在错误信息
		if _, exists := controllers.UsersLoginInfo[token]; !exists {
			c.JSON(http.StatusOK, controllers.Response{StatusCode: 1, StatusMsg: "token鉴权失败, 非法操作"})
			c.Abort()
			return
		}
		// 如果token存在，将登录用户的ID存储在上下文中，并使用c.Next()继续请求处理流程。
		c.Set("userID", controllers.UsersLoginInfo[token].ID)
		c.Next()
	}
}
