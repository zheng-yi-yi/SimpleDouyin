package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
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

		if _, exists := response.UsersLoginInfo[token]; !exists {
			c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 1, StatusMsg: authFailedMsg})
			c.Abort()
			return
		}

		c.Set(userIDKey, response.UsersLoginInfo[token].ID)
		c.Next()
	}
}
