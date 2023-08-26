package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers"
)

const (
	userIDKey     = "userID"
	authFailedMsg = "Token鉴权失败，非法操作"
)

// JWTAuth is a middleware for simple JWT authentication.
// It specifies where to get the token from (query or form).
func JWTAuth(where string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		if where == "form" {
			token = c.PostForm("token")
		} else {
			token = c.Query("token")
		}

		if _, exists := controllers.UsersLoginInfo[token]; !exists {
			c.JSON(http.StatusUnauthorized, controllers.Response{StatusCode: 1, StatusMsg: authFailedMsg})
			c.Abort()
			return
		}

		c.Set(userIDKey, controllers.UsersLoginInfo[token].ID)
		c.Next()
	}
}
