package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

// 密钥
var JwtKey = []byte(config.AUTH_KEY)

// 一些常量
const (
	ISSUE            = "SimpleDouyin"
	CheckFailed      = "校验失败"
	ParseTokenFailed = "解析token失败"
	GetTokenFailed   = "获取token失败"
)

// MyClaims，定义 JWT 的声明
type MyClaims struct {
	UserID   uint
	Username string
	Password string
	jwt.RegisteredClaims
}

// 鉴权失败响应
type AuthFailResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// GenerateToken , 生成token
func GenerateToken(userId uint, username, password string) (string, error) {
	claims := MyClaims{
		userId,
		username, // 用户名
		password, // 密码
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间（通常相对于当前时间设置）
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 令牌的发行时间（即创建时间）
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 令牌生效时间（在此时间之前令牌无效）
			Issuer:    ISSUE,                                              // 创建此令牌的实体
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Auth , 验证 JWT token，并将其中的用户信息解析后存储在 Gin 的上下文中
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		tokenString := c.Query("token")
		if len(tokenString) == 0 {
			c.Set("userID", uint(0))
			c.Next()
			return
		}
		// 解析 token
		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, AuthFailResponse{StatusCode: 1, StatusMsg: ParseTokenFailed})
			c.Abort()
			return
		}
		// 校验鉴权的声明
		if token != nil {
			claims, ok := token.Claims.(*MyClaims)
			if ok && token.Valid {
				c.Set("userID", claims.UserID)
				c.Next()
				return
			}
		}
		c.JSON(http.StatusInternalServerError, AuthFailResponse{StatusCode: 1, StatusMsg: CheckFailed})
		c.Abort()
	}
}

// UserPublishAuth , 用户发布视频前进行用户验证
// 由于客户端提交表单数据时，
// 包含的token字段是在请求体中以 multipart/form-data 的方式进行传递
// 因此需要使用 c.PostForm 方法来获取表单数据中的字段值
func UserPublishAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		tokenString := c.PostForm("token")
		if len(tokenString) == 0 {
			c.Set("userID", uint(0))
			c.Next()
			return
		}
		// 解析 token
		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			log.Println("token解析失败")
			c.JSON(http.StatusBadRequest, AuthFailResponse{StatusCode: 1, StatusMsg: ParseTokenFailed})
			c.Abort()
			return
		}
		// 校验鉴权的声明
		if token != nil {
			claims, ok := token.Claims.(*MyClaims)
			if ok && token.Valid {
				c.Set("userID", claims.UserID)
				c.Next()
				return
			}
		}
		log.Println("校验鉴权失败")
		c.JSON(http.StatusBadRequest, AuthFailResponse{StatusCode: 1, StatusMsg: CheckFailed})
		c.Abort()
	}
}
