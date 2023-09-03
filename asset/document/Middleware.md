# Middleware

本篇文档记录JWT（JSON Web Token）中间件的设计思路和实现过程。

---

## JWT 😊

`JWT` 是一种用于在网络上安全传输信息的开放标准（RFC 7519），通常用于身份验证和信息交换。在本文中，我们将使用 `Golang` 和 `Gin` 框架来实现 `JWT` 中间件，用于验证用户的身份并将其信息存储在 `Gin` 的上下文中。

## 设计思路 💡

### （1）密钥管理

首先，我们需要定义一个密钥（JwtKey）来对JWT进行签名和验证。密钥应该是一个安全的字节序列，并且应该妥善保管。

```go
var JwtKey = []byte("douyin")
```

### （2）JWT声明定义

我们需要定义 `JWT` 的声明，以便在生成和解析 `JWT` 时使用。在本项目中，我定义了 `MyClaims` 结构体，其中包括用户ID、用户名、密码以及一些标准的 `JWT` 声明（如过期时间、发行时间、发行者）。

```go
type MyClaims struct {
	UserID   uint
	Username string
	Password string
	jwt.RegisteredClaims
}
```

### （3）生成JWT Token

要生成 `JWT Token`，我们使用 `GenerateToken` 函数，该函数接受用户ID、用户名和密码作为参数，并返回一个签名后的`JWT Token`。在生成过程中，我们设置了`JWT`的一些标准声明，如过期时间和发行者。

```go
func GenerateToken(userId uint, username, password string) (string, error) {
	claims := MyClaims{
		userId,
		username,
		password, 
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()), 
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "SimpleDouyin",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
```

### （4）鉴权中间件

我们定义了一个 `Auth` 中间件函数，用于验证 `JWT Token` 并将用户信息存储在 `Gin` 的上下文中。

该中间件首先获取请求中的 `Token`，然后尝试解析它。如果解析成功且 `Token` 有效，我们将 `用户ID` 存储在上下文中，以便后续处理函数使用。如果解析失败或 `Token无效` ，我们将返回鉴权失败的响应。

```go
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		tokenString := c.Query("token")
		if len(tokenString) == 0 {
			c.JSON(http.StatusBadRequest, AuthFailResponse{StatusCode: 1, StatusMsg: "无效token"})
			c.Abort()
			return
		}
		// 解析 token
		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, AuthFailResponse{StatusCode: 1, StatusMsg: "token解析失败"})
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
		c.JSON(http.StatusBadRequest, AuthFailResponse{StatusCode: 1, StatusMsg: "校验失败"})
		c.Abort()
	}
}
```

### （5）用户发布视频前的鉴权

对于用户发布视频的鉴权，我们定义了 `UserPublishAuth` 中间件函数。

由于客户端提交表单数据时，`Token`字段是以 `multipart/form-data` 方式传递的，因此我们需要使用 `c.PostForm` 方法来获取 `Token` 值。鉴权逻辑与上述 `Auth` 中间件类似。

```go
func UserPublishAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		tokenString := c.PostForm("token")
		if len(tokenString) == 0 {
			c.JSON(http.StatusBadRequest, AuthFailResponse{StatusCode: 1, StatusMsg: "无效token"})
			c.Abort()
			return
		}
		// 解析 token
		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, AuthFailResponse{StatusCode: 1, StatusMsg: "token解析失败"})
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
		c.JSON(http.StatusBadRequest, AuthFailResponse{StatusCode: 1, StatusMsg: "校验失败"})
		c.Abort()
	}
}
```

---

## 如何使用

当我们需要生成token时，可以这样：

```go
// 生成Token
token, err := GenerateToken(userId, username, password)
if err != nil {
    // 处理错误
} else {
    // 将Token发送给客户端
}
```

用户鉴权的话，举一个例子：

```go
router.POST("/publish/action/", UserPublishAuth(), func(c *gin.Context) {
    // 从上下文中获取用户ID
    userID := c.Value("userID")
    // 处理视频发布逻辑
    // ...
})
```

---

## 参考资料

了解更多关于JWT的信息，请参考以下文档：

- [JWT 官方文档](https://jwt.io/introduction)
- [jwt-go 包文档](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)
- [RFC 7519 - JSON Web Token (JWT)](https://datatracker.ietf.org/doc/html/rfc7519)